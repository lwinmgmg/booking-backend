package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	connectDB()
}

func connectDB() {
	var err error
	dsn := "host=localhost user=lwinmgmg password=letmein dbname=booking_system port=5432 sslmode=disable TimeZone=Asia/Rangoon"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		connectDB()
	}
	return db
}
