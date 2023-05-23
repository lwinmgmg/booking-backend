package models

import "gorm.io/gorm"

type BookingLine struct {
	gorm.Model
	BookingID  uint
	Booking    Booking `gorm:"foreignKey:BookingID"`
	ProductID  uint
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   float64
	UnitID     uint
	Unit       Unit `gorm:"foreignKey:UnitID"`
	TotalPrice float64
}

func (user BookingLine) TableName() string {
	return "booking_line"
}
