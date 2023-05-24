package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/models"
	"github.com/lwinmgmg/booking-backend/schemas"
)

type ProductController struct {
}

func (pc *ProductController) GetItems(c *gin.Context) {
	products := []models.Product{}
	if err := db.Model(&models.Product{}).Limit(10).Offset(0).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	responseProductList := []schemas.ProductResponse{}
	for _, product := range products {
		if product.UnitID > 0 {
			if err := db.First(&(product.Unit), product.UnitID).Error; err != nil {
				c.JSON(http.StatusBadRequest, map[string]any{
					"success": false,
					"message": err.Error(),
				})
				return
			}
		}
		prodTemp := schemas.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Type:     product.Type,
			Price:    product.SalePrice,
			UomID:    product.Unit.ID,
			UomName:  product.Unit.Name,
			ImageUrl: product.ImageUrl,
		}
		responseProductList = append(responseProductList, prodTemp)
	}
	c.JSON(http.StatusOK, responseProductList)

}
