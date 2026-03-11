package main

import (
	"os"
	"zhix-backend/config"
	"zhix-backend/middleware"
	"zhix-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
	}
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
	} else {
	}
}

func main() {
	middleware.InitJWT()
	config.InitDB()
	config.InitRedis()

	r := gin.Default()
	r.Use(corsMiddleware())

	routes.SetupRoutes(r)

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
