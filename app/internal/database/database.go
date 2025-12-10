package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB nesnesini global olarak tanımlıyoruz ki diğer paketlerden erişebilelim
var DB *gorm.DB

func ConnectDB() {
	// .env dosyasından okuduğumuz değerlerle DSN (Data Source Name) oluşturuyoruz
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı! Hata: ", err)
	}

	log.Println("Veritabanı bağlantısı başarıyla sağlandı! 🚀")
}
