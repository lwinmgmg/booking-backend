package schemas

import (
	"errors"
)

type BookingLineCreate struct {
	ProductID *uint `json:"product_id"`
	Quantity  *float64
}

func (bl *BookingLineCreate) Validate() error {
	return nil
}

type BookingCreate struct {
	BookingLines []*BookingLineCreate `json:"booking_lines"`
}

func (bc *BookingCreate) Validate() error {
	if len(bc.BookingLines) == 0 {
		return errors.New("at least one booking is required")
	}
	for _, bl := range bc.BookingLines {
		if err := bl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookingReport struct {
	PartnerID string `json:"partner_id"`
	Date      string `json:"date"`
}

type BookingReportCsv struct {
	CustomerName string
	Phone        string
	Nrc          string
	ProductName  string
	Type         string
	SalePrice    float64
	Quantity     float64
	Unit         string
	TotalPrice   float64
	CreatedAt    string
}
