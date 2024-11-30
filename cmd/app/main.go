package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sijiaoh/golang-web/internal/router"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run()
}
