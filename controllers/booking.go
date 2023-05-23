package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/models"
	"github.com/lwinmgmg/booking-backend/schemas"
	"gorm.io/gorm"
)

type BookingController struct {
}

func (bc *BookingController) Book(c *gin.Context) {
	userID := 8
	bookingCreate := &schemas.BookingCreate{}
	if err := c.ShouldBind(bookingCreate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	booking := &models.Booking{
		PartnerID: *bookingCreate.PartnerID,
		UserID:    uint(userID),
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(booking).Error; err != nil {
			return err
		}
		for _, bl := range bookingCreate.BookingLines {
			product := &models.Product{}
			if err := db.First(product, *bl.ProductID).Error; err != nil {
				return err
			}
			bookingLine := &models.BookingLine{
				ProductID: product.ID,
				BookingID: booking.ID,
				Quantity:  *bl.Quantity,
				UnitID:    product.UnitID,
			}
			if err := tx.Create(bookingLine).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	db.Model(&models.BookingLine{}).Where("booking_id = ?", booking.ID).Find(&booking.BookingLines)
	c.JSON(http.StatusOK, booking)
}
