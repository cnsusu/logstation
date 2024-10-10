package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//释放资源
		defer func() {
			if err := recover(); err != nil {
				log.Printf("err=%s, stack=%s", fmt.Sprint(err), string(debug.Stack()))
				_ = c.AbortWithError(999, errors.New(fmt.Sprintf("%s", err)))
			}
		}()

		c.Next()
	}
}
