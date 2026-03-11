package controllers

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const ffmpegDir = "/tmp/zhix-ffmpeg"

var ffmpegPathMu sync.Mutex

// ffmpegBinary returns the path to a working ffmpeg binary.
// Checks PATH and ffmpegDir first; downloads only if not found.
// Safe for concurrent calls.
func ffmpegBinary() (string, error) {
	ffmpegPathMu.Lock()
	defer ffmpegPathMu.Unlock()

	// 1. System PATH
	if p, err := exec.LookPath("ffmpeg"); err == nil {
		return p, nil
	}

	// 2. Previously downloaded binary
	local := filepath.Join(ffmpegDir, "ffmpeg")
	if runtime.GOOS == "windows" {
		local += ".exe"
	}
	if _, err := os.Stat(local); err == nil {
		return local, nil
	}

	// 3. Download (only reaches here when binary is absent)
	if err := downloadFFmpeg(); err != nil {
		return "", fmt.Errorf("ffmpeg not found and download failed: %w", err)
	}
	if _, err := os.Stat(local); err != nil {
		return "", fmt.Errorf("ffmpeg binary not available after download")
	}
	return local, nil
}

func downloadFFmpeg() error {
	if err := os.MkdirAll(ffmpegDir, 0755); err != nil {
		return err
	}

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	// BtbN static builds: https://github.com/BtbN/FFmpeg-Builds/releases
	var downloadURL, archiveName string
	switch {
	case goos == "linux" && goarch == "amd64":
		archiveName = "ffmpeg-master-latest-linux64-gpl.tar.xz"
	case goos == "linux" && goarch == "arm64":
		archiveName = "ffmpeg-master-latest-linuxarm64-gpl.tar.xz"
	case goos == "darwin":
		// macOS: use evermeet.cx static build
		archiveName = "ffmpeg.zip"
		downloadURL = "https://evermeet.cx/ffmpeg/getrelease/zip"
	case goos == "windows":
		archiveName = "ffmpeg-master-latest-win64-gpl.zip"
	default:
		return fmt.Errorf("unsupported platform: %s/%s", goos, goarch)
	}

	if downloadURL == "" {
		downloadURL = "https://github.com/BtbN/FFmpeg-Builds/releases/download/latest/" + archiveName
	}

	archivePath := filepath.Join(ffmpegDir, archiveName)
	if err := downloadFile(archivePath, downloadURL); err != nil {
		return fmt.Errorf("download ffmpeg: %w", err)
	}
	defer os.Remove(archivePath)

	return extractFFmpeg(archivePath, ffmpegDir, goos)
}

// proxyClient returns an http.Client that honours HTTPS_PROXY / HTTP_PROXY / NO_PROXY
// environment variables, identical to curl/wget proxy behaviour.
func proxyClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return http.ProxyFromEnvironment(req)
			},
		},
	}
}

func downloadFile(dest, rawURL string) error {
	client := proxyClient()
	resp, err := client.Get(rawURL) //nolint:gosec
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d fetching %s", resp.StatusCode, rawURL)
	}
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func extractFFmpeg(archivePath, destDir, goos string) error {
	binName := "ffmpeg"
	if goos == "windows" {
		binName = "ffmpeg.exe"
	}
	destBin := filepath.Join(destDir, binName)

	if strings.HasSuffix(archivePath, ".zip") {
		return extractFromZip(archivePath, binName, destBin)
	}
	// .tar.xz / .tar.gz
	return extractFromTarXZ(archivePath, binName, destBin)
}

func extractFromZip(archivePath, binName, destBin string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		if filepath.Base(f.Name) == binName {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()
			return writeExecutable(destBin, rc)
		}
	}
	return fmt.Errorf("%s not found in zip", binName)
}

func extractFromTarXZ(archivePath, binName, destBin string) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Try gzip first, then raw tar (xz not in stdlib — handle via pre-decompressed or pipe)
	var tr *tar.Reader
	gr, err := gzip.NewReader(f)
	if err == nil {
		tr = tar.NewReader(gr)
		defer gr.Close()
	} else {
		// xz: decompress via system xz or xzcat if available
		f.Seek(0, 0)
		xzCmd := exec.Command("xzcat", archivePath)
		xzOut, err2 := xzCmd.StdoutPipe()
		if err2 != nil {
			return fmt.Errorf("xz not available and gzip failed: %w", err)
		}
		if err2 = xzCmd.Start(); err2 != nil {
			return fmt.Errorf("xzcat start: %w", err2)
		}
		tr = tar.NewReader(xzOut)
		defer xzCmd.Wait()
	}

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if filepath.Base(hdr.Name) == binName && hdr.Typeflag == tar.TypeReg {
			return writeExecutable(destBin, tr)
		}
	}
	return fmt.Errorf("%s not found in archive", binName)
}

func writeExecutable(dest string, r io.Reader) error {
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

// ─── Transcode Handler ────────────────────────────────────────────────────────

// TranscodeToFMP4 converts a plain MP4 (or WebM) to fragmented MP4 suitable for MSE.
// Request: JSON { "url": "https://..." }  OR multipart file upload (field "file")
// Response: fMP4 binary stream with Content-Type video/mp4
func TranscodeToFMP4(c *gin.Context) {
	ffmpeg, err := ffmpegBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ── Resolve input ──────────────────────────────────────────────────────
	tmpDir, err := os.MkdirTemp("", "zhix-transcode-*")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tmp dir: " + err.Error()})
		return
	}
	defer os.RemoveAll(tmpDir)

	inputPath := filepath.Join(tmpDir, "input.mp4")

	if videoURL := c.PostForm("url"); videoURL != "" {
		// Download from URL
		if err := downloadFile(inputPath, videoURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "download input: " + err.Error()})
			return
		}
	} else {
		// Multipart upload
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "provide 'url' or 'file'"})
			return
		}
		if err := c.SaveUploadedFile(file, inputPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "save upload: " + err.Error()})
			return
		}
	}

	// ── Run ffmpeg ─────────────────────────────────────────────────────────
	outputPath := filepath.Join(tmpDir, "output.mp4")
	cmd := exec.Command(ffmpeg,
		"-i", inputPath,
		"-an",
		"-movflags", "frag_keyframe+empty_moov+faststart",
		"-c:v", "copy",
		"-y",
		outputPath,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "ffmpeg failed: " + err.Error(),
			"detail": string(out),
		})
		return
	}

	// ── Stream result ──────────────────────────────────────────────────────
	c.Header("Content-Disposition", "attachment; filename=\"output_fmp4.mp4\"")
	c.File(outputPath)
}
