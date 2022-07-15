package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StatusLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		log.Print(latency)

		status := ctx.Writer.Status()

		log.Println(status)

	}

}
