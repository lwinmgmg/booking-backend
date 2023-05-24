package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string
	Type      string
	Quantity  float64
	SalePrice float64
	UnitID    uint
	Unit      Unit `gorm:"foreignKey:UnitID"`
	ImageUrl  string
}

func (user Product) TableName() string {
	return "product"
}
