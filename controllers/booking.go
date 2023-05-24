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
	userID, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user := &models.User{}
	if err := db.First(user, userID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": err.Error(),
		})
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
		PartnerID: user.PartnerID,
		UserID:    userID,
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
				ProductID:  product.ID,
				BookingID:  booking.ID,
				Quantity:   *bl.Quantity,
				UnitID:     product.UnitID,
				TotalPrice: product.SalePrice * (*bl.Quantity),
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
	dataList := make([]schemas.BookingReportCsv, 0)
	if partner_id == "" {
		db.Raw(fmt.Sprintf(`
		SELECT rp.first_name || ' ' || rp.last_name AS customer_name, rp.phone AS phone, rp.nrc, pt.name AS product_name, pt.type, pt.sale_price, bl.quantity, uom.name AS unit, bl.total_price, bk.created_at FROM booking_line AS bl
		INNER JOIN booking bk ON bk.id=bl.booking_id
		INNER JOIN product pt ON pt.id=bl.product_id
		LEFT JOIN uom ON uom.id=bl.unit_id
		LEFT JOIN res_partner rp ON rp.id=bk.partner_id
		WHERE bk.created_at::DATE = ?
		`), dateStr).Scan(&dataList)
	} else {
		db.Raw(fmt.Sprintf(`
		SELECT rp.first_name || ' ' || rp.last_name AS customer_name, rp.phone AS phone, rp.nrc, pt.name AS product_name, pt.type, pt.sale_price, bl.quantity, uom.name AS unit, bl.total_price, bk.created_at FROM booking_line AS bl
		INNER JOIN booking bk ON bk.id=bl.booking_id
		INNER JOIN product pt ON pt.id=bl.product_id
		LEFT JOIN uom ON uom.id=bl.unit_id
		LEFT JOIN res_partner rp ON rp.id=bk.partner_id
		WHERE bk.created_at::DATE = ? AND bk.partner_id = ?
		`), dateStr, partner_id).Scan(&dataList)
	}
	filename := "assets/booking_report.csv"
	file, err := os.Create(filename)
	csvWriter := csv.NewWriter(file)
	records := [][]string{}
	records = append(records, []string{
		"Customer",
		"Phone",
		"NRC",
		"Product",
		"Type",
		"SalePrice",
		"Quantity",
		"Unit",
		"Total Price",
		"Order Date",
	})
	defer csvWriter.Flush()
	for _, data := range dataList {
		records = append(records, []string{
			data.CustomerName,
			data.Phone,
			data.Nrc,
			data.ProductName,
			data.Type,
			fmt.Sprint(data.SalePrice),
			fmt.Sprint(data.Quantity),
			data.Unit,
			fmt.Sprint(data.TotalPrice),
			data.CreatedAt,
		})
	}
	fmt.Println(records)
	csvWriter.WriteAll(records)
	file.Close()
	c.File(filename)
	os.Remove(filename)
}
