package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("Uyarı: .env dosyası bulunamadı, ortam değişkenleri bekleniyor.")
	}

	// 2. Gin sunucusunu başlat
	r := gin.Default()

	// 3. Basit bir test rotası (Health Check)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "success",
		})
	})

	// 4. Sunucuyu ayağa kaldır
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Sunucu %s portunda çalışıyor...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Sunucu başlatılamadı:", err)
	}
}
