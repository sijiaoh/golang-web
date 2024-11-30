package router

import (
	"github.com/gin-gonic/gin"
	controllers_v1 "github.com/sijiaoh/golang-web/internal/controllers/v1"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", controllers_v1.PingShow)
}
