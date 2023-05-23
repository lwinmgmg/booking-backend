package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/api"
)

func main() {
	r := gin.Default()
	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"POST", "PUT", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)
	v1 := r.Group("/api/v1")
	routerApi := api.NewApiRouter(v1)
	routerApi.Register()
	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
