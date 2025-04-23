package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		logger := zap.L().With(
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("ip", c.ClientIP()),
		)

		if len(c.Errors) > 0 {
			logger.Error("Request completed with errors",
				zap.Strings("errors", c.Errors.Errors()),
			)
		} else {
			logger.Info("Request completed")
		}
	}
}
