package routers

import (
	"github.com/gin-gonic/gin"
	"logstation/client"
	"logstation/controllers"
	"logstation/middleware"
	"net/http"
)

func InitRouter(router *gin.Engine, hub *client.Hub) {
	router.LoadHTMLGlob("views/*")
	router.Static("/assets", "./assets")
	router.GET("/login", controllers.HandleLogin)
	router.GET("/", middleware.Middleware(), controllers.HandleIndex)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(hub, w, r)
	})
}
