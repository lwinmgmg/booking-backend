package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/controllers"
	"github.com/lwinmgmg/booking-backend/models"
	"github.com/lwinmgmg/booking-backend/services"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB = services.GetDB()
	User          = models.User{}
)

type ApiRouter struct {
	Router *gin.RouterGroup
}

func (apiRouter *ApiRouter) Register() error {
	// User Routes
	userController := &controllers.UserController{}
	apiRouter.Router.POST("/user", userController.CreateUser)
	apiRouter.Router.POST("/login", userController.Login)

	bookingController := &controllers.BookingController{}
	apiRouter.Router.POST("/bookings", bookingController.Book)
	apiRouter.Router.GET("/report/bookings", bookingController.BookingDetailReport)
	return nil
}

func NewApiRouter(router *gin.RouterGroup) *ApiRouter {
	return &ApiRouter{
		Router: router,
	}
}
