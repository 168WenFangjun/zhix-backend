package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var baseURL string = "https://raw.githubusercontent.com/168WenFangjun/zhix-articles/main"

var coverImages = []string{
	baseURL + "/images/1.png",
	baseURL + "/images/2.webp",
	baseURL + "/images/3.png",
	baseURL + "/images/4.webp",
	baseURL + "/images/5.png",
}

var coverCartoons = []string{
	baseURL + "/cartoons/cat.webp",
	baseURL + "/cartoons/cold.webp",
}

type videoEntry struct {
	Video string
	Audio string
}

var coverVideos = []videoEntry{
	{baseURL + "/videos/cat_fmp4.mp4", baseURL + "/videos/cat.mp3"},
	{baseURL + "/videos/cold_fmp4.mp4", baseURL + "/videos/cold.mp3"},
	{baseURL + "/videos/flower_fmp4.mp4", baseURL + "/videos/flower.mp3"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetCoverImage(c *gin.Context) {
	randomIndex := rand.Intn(len(coverImages))
	c.JSON(http.StatusOK, gin.H{
		"coverImage": coverImages[randomIndex],
	})
}

func GetCoverCartoon(c *gin.Context) {
	randomIndex := rand.Intn(len(coverCartoons))
	c.JSON(http.StatusOK, gin.H{
		"coverImage": coverCartoons[randomIndex],
	})
}

func GetCoverVideo(c *gin.Context) {
	e := coverVideos[rand.Intn(len(coverVideos))]
	c.JSON(http.StatusOK, gin.H{
		"coverVideo": e.Video,
		"coverAudio": e.Audio,
	})
}
