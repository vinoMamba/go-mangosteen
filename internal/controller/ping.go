package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandle(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
