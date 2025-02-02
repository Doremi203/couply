package app

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func newGinRouter(
	log *slog.Logger,
) *gin.Engine {
	router := gin.New()

	router.Use(customLoggerMiddleware(log))
	router.Use(gin.Recovery())

	return router
}

func customLoggerMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		reqTime := time.Since(start)
		statusCode := c.Writer.Status()
		log.Info(
			"request",
			"path", c.Request.URL.Path,
			"status", statusCode,
			"processing time", reqTime,
		)
	}
}
