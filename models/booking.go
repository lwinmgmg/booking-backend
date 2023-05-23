package models

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	PartnerID    uint
	Partner      Partner `gorm:"foreignKey:PartnerID"`
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	BookingLines []BookingLine
}

func (user Booking) TableName() string {
	return "booking"
}
