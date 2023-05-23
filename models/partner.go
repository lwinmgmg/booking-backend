package models

import "gorm.io/gorm"

type Partner struct{
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Nrc       string
	Address   string
}

func (user Partner) TableName() string {
	return "res_partner"
}
