package models

import (
	"github.com/lwinmgmg/booking-backend/services"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = services.GetDB()
	db.AutoMigrate(&Partner{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Unit{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&BookingLine{})
	db.AutoMigrate(&Booking{})
}
