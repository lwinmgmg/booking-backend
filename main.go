package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/booking-backend/api"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	routerApi := api.NewApiRouter(v1)
	routerApi.Register()
	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
