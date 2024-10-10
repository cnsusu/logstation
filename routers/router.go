package routers

import (
	"github.com/gin-gonic/gin"
	"logstation/controllers"
	"logstation/middleware"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.Static("/assets", "./assets")
	router.GET("/login", controllers.HandleLogin)
	router.GET("/", middleware.Middleware(), controllers.HandleIndex)
}
