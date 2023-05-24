package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/services"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = services.GetDB()
)

const (
	UserHeader = "user_id"
)

type Controller struct {
	UserID uint
}

func getUserId(c *gin.Context) (uint, error) {
	userIDStr := c.GetHeader(UserHeader)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": "Not authenticated",
		})
	}
	return uint(userID), err
}
