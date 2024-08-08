package middlewares

import (
	"gorestAPI/db"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DBConnectionCheck(db db.DBinterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !db.ConnectionCheck() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
			return
		}
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.New()
		file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			logger.Info("Failed to log to file, using default stderr")
		}
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request
		logger.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"status":   c.Writer.Status(),
			"duration": time.Since(start),
			"path":     c.Request.URL.Path,
			"message":  c.MustGet("msg"),
		}).Info("Request processed")
	}
}
