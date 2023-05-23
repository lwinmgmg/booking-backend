package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	PartnerID   string
	BookingTime time.Time
}

func (user Booking) TableName() string {
	return "booking"
}
