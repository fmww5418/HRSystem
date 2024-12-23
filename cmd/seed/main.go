package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"HRSystem/config"
	"HRSystem/migrate/seeds"
	libdb "HRSystem/src/lib/db"
)

func main() {
	filepath := "config/local.json"
	config.LoadConfig(&filepath)

	dsn := libdb.GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	seeds.Case1(db)
}
