package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/models"
	"github.com/lwinmgmg/booking-backend/schemas"
	"gorm.io/gorm"
)

type BookingController struct {
}

func (bc *BookingController) Book(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
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
		UserID:    userId,
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

func (bc *BookingController) BookingDetailReport(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	partner_id := c.Query("partner_id")
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": "date can't be empty",
		})
		return
	}
	_, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	data := make([]map[string]any, 0)
	if partner_id == "" {
		db.Raw(fmt.Sprintf(`
		SELECT * FROM booking_line AS bl
		INNER JOIN booking bk ON bk.id=bl.booking_id
		LEFT JOIN partner pt ON pt.id=bk.partner_id
		WHERE bk.created_at::DATE = ?
		`), dateStr).Scan(&data)
	} else {
		db.Raw(fmt.Sprintf(`
		SELECT * FROM booking_line AS bl
		INNER JOIN booking bk ON bk.id=bl.booking_id
		LEFT JOIN partner pt ON pt.id=bk.partner_id
		WHERE bk.created_at::DATE = ? AND bk.partner_id = ?
		`), dateStr, partner_id).Scan(&data)
	}

	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}
	file, err := os.Create("assets/abc.csv")
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	csvWriter.WriteAll(records)
	file.Close()
	c.File("assets/abc.csv")
	os.Remove("assets/abc.csv")
}
