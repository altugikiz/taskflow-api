package main

import (
	"log"
	"os"

	"github.com/altugikiz/app/internal/database" // Burayı kendi modül ismine göre ayarla
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Ortam değişkenlerini yükle
	if err := godotenv.Load(); err != nil {
		log.Println(".env dosyası bulunamadı.")
	}

	// 2. Veritabanına bağlan
	database.ConnectDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
    
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}