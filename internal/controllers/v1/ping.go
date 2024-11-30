package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingShow(c *gin.Context) {
	c.HTML(http.StatusOK, "ping/show", gin.H{
		"title": "ping",
	})
}
