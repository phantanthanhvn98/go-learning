package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	dsn := "host=ep-billowing-sunset-a5pok1yc-pooler.us-east-2.aws.neon.tech user=tax_owner password=b2GUjBaR3iSo dbname=tax port=5432 sslmode=require"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	log.Println("Connected to database")
}
