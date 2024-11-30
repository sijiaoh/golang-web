package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sijiaoh/golang-web/internal/router"
	"github.com/sijiaoh/golang-web/internal/views"
)

func main() {
	r := gin.Default()
	r.HTMLRender = views.CreateViews()
	router.SetupRouter(r)
	r.Run()
}
