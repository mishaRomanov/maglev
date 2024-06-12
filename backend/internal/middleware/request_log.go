package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		log.Infof("New request path:%s method:%s time:%s ",
			ctx.Request.URL,
			ctx.Request.Method,
			time.Now().Format(time.TimeOnly))
	}
}
