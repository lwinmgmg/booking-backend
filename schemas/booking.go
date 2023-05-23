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
	PartnerID    *uint                `json:"partner_id"`
	BookingLines []*BookingLineCreate `json:"booking_lines"`
}

func (bc *BookingCreate) Validate() error {
	if bc.PartnerID == nil {
		return errors.New("partner can't be empty")
	}
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
