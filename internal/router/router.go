package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/go-mangosteen/internal/controller"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
	//404 路由
	r.NoRoute(func(c *gin.Context) {
		acceptStr := c.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			c.String(http.StatusNotFound, "页面404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"err_code":    1,
				"err_message": "未定义的路由",
			})
		}

	})

	r.GET("/ping", controller.PingHandle)
	return r
}
