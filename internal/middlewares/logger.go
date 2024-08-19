package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func ZerologMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		logger := zerolog.New(os.Stdout).With().
			Timestamp().
			Str("request_id", ginContext.GetString("X-Request-Id")). // Example for request ID if available
			Logger()

		timeSince := time.Now()

		ginContext.Set("logger", logger)

		ginContext.Next()

		latency := time.Since(timeSince)
		status := ginContext.Writer.Status()
		logger.Info().Msgf("Request completed in %v | Status: %d", latency, status)
	}
}

func LoggerFromContext(c *gin.Context) zerolog.Logger {
	logger, err := c.MustGet("logger").(zerolog.Logger)
	if !err {
		panic("logger not found in context")
	}

	return logger
}
