package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/go-mangosteen/internal/controller"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.PingHandle)
	return r
}
