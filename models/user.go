package models

import "gorm.io/gorm"

type UserRole string

const (
	ADMIN  UserRole = "ADMIN"
	PORTAL UserRole = "PORTAL"
)

type User struct {
	gorm.Model
	Role      UserRole
	PartnerID uint    `gorm:"unique"`
	Partner   Partner `gorm:"foreignKey:PartnerID"`
	Username  string  `gorm:"unique"`
	Password  string
}

func (user User) TableName() string {
	return "res_users"
}
