package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"CipherX/config"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// Start time
		startTime := time.Now()

		b, _ := c.Copy().GetRawData()

		c.Request.Body = io.NopCloser(bytes.NewReader(b))

		// Process request
		c.Next()

		// End time
		endTime := time.Now()

		config.GinLOG.Info("Request Response",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("request_time", TimeFormat(startTime)),
			zap.String("response_time", TimeFormat(endTime)),
			zap.String("cost_time", endTime.Sub(startTime).String()),
		)
	}
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
