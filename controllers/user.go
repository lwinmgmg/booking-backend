package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/models"
	"github.com/lwinmgmg/booking-backend/schemas"
	"gorm.io/gorm"
)

type UserController struct {
	Controller
}

func (userApi *UserController) CreateUser(c *gin.Context) {
	userCreate := &schemas.UserCreate{}
	userRole := models.PORTAL
	if err := c.ShouldBind(userCreate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if err := userCreate.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	partner := &models.Partner{
		FirstName: *userCreate.FirstName,
		LastName:  *userCreate.LastName,
		Email:     *userCreate.Email,
		Phone:     *userCreate.Phone,
		Nrc:       *userCreate.Nrc,
		Address:   *userCreate.Address,
	}
	user := &models.User{
		Username:  *userCreate.Username,
		Password:  *userCreate.Password,
		PartnerID: partner.ID,
		Role:      userRole,
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(partner).Error; err != nil {
			return err
		}
		user.PartnerID = partner.ID
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (userController *UserController) Login(c *gin.Context) {
	userLogin := &schemas.UserLogin{}
	if err := c.ShouldBind(userLogin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if err := userLogin.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user := &models.User{}
	if err := db.Model(user).Where("username = ?", *userLogin.UserName).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]any{
			"success": false,
			"message": fmt.Sprintf("User[%v] Not Found", *userLogin.UserName),
		})
		return
	}
	if user.Password != *userLogin.Password {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": "Wrong Password",
		})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, schemas.UserLoginResponse{
		Username: user.Username,
		ID:       user.ID,
		Role:     string(user.Role),
	})
}

func (userApi *UserController) DeleteUser(c *gin.Context) {
}
