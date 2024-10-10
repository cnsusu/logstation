package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func HandleIndex(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err=%s, stack=%s", fmt.Sprint(err), string(debug.Stack()))
		}
	}()

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
