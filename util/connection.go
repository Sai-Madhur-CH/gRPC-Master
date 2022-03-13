package util

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Sai-Madhur-CH/gRPC-Master/models"
)

var db *gorm.DB

// Connect -
func Connect() *gorm.DB {
	var e error
	dbPath := "host=127.0.0.1 port=5430 dbname=chatDB sslmode=disable user=postgres password=postgres"
	db, e = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbPath,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if e != nil {
		panic("failed to connect to database")
	}

	fmt.Println("DB Connected...")
	autoMigrate()
	return db
}

func autoMigrate() {
	fmt.Print("Started Auto-Migrations...")
	db.AutoMigrate(&models.User{})
	fmt.Print("Done with Auto-Migrations.")
}
