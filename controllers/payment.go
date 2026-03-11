package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProcessUnionPayH5 处理银联UPOP H5支付
func ProcessUnionPayH5(c *gin.Context) {
	var input struct {
		OrderId string `json:"orderId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.OrderId == "" {
		input.OrderId = generatePaymentID()
	}

	// TODO: 集成真实银联SDK，需要商户号、证书和签名
	// 当前返回模拟支付页面
	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>银联支付</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; background: #f5f5f5; }
        .container { max-width: 400px; margin: 50px auto; background: white; padding: 30px; border-radius: 12px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h2 { color: #333; text-align: center; }
        .info { margin: 20px 0; padding: 15px; background: #f9f9f9; border-radius: 8px; }
        .info p { margin: 8px 0; color: #666; }
        .btn { width: 100%; padding: 15px; background: #e60012; color: white; border: none; border-radius: 8px; font-size: 16px; cursor: pointer; margin-top: 10px; }
        .btn:hover { background: #cc0010; }
        .btn-secondary { background: #6b7280; }
        .btn-secondary:hover { background: #4b5563; }
        .note { text-align: center; color: #999; font-size: 12px; margin-top: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <h2>🏦 银联UPOP支付</h2>
        <div class="info">
            <p><strong>订单号：</strong>` + input.OrderId + `</p>
            <p><strong>支付金额：</strong>¥39.99</p>
            <p><strong>商品：</strong>极志社区会员（1个月）</p>
        </div>
        <button class="btn" onclick="handlePay(true)">确认支付</button>
        <button class="btn btn-secondary" onclick="handlePay(false)">模拟支付失败</button>
        <p class="note">这是模拟支付页面<br/>实际使用需配置银联商户号和证书</p>
    </div>
    <script>
        function handlePay(success) {
            if (success) {
                window.location.href = '/payment/success?orderId=` + input.OrderId + `';
            } else {
                window.location.href = '/payment/failure?orderId=` + input.OrderId + `';
            }
        }
    </script>
</body>
</html>`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func generatePaymentID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return "pay_" + hex.EncodeToString(b)
}
