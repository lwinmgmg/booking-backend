package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name string
}

func (user Unit) TableName() string {
	return "uom"
}
