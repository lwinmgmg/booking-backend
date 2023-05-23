package controllers

import (
	"github.com/lwinmgmg/booking-backend/services"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = services.GetDB()
)

type Controller struct {
	UserID uint
}
