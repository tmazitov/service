package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContentLimiter(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a limited reader that will stop reading after maxBytes
		reader := io.LimitReader(c.Request.Body, maxBytes)
		// Read the body into a buffer
		body, err := io.ReadAll(reader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
			c.Abort()
			return
		}

		// Check if the reader has more data than maxBytes
		if _, err := io.Copy(io.Discard, c.Request.Body); err == nil {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "request body too large"})
			c.Abort()
			return
		}

		// Replace the request body with the read body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		c.Next()
	}
}
