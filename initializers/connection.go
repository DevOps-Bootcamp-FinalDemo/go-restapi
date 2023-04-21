package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	DSN := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database", err)
	} else {
		log.Println("DB Connected")
	}
}
