// Package: middleware
// Description: Provides middleware functionalities for the API, including token validation
// and integration with SwingClient.

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// RemoveJSONSuffixFromParams is a Gin middleware that removes the ".json" suffix
// from request parameters (both key and value).
// This ensures that API routes can handle requests with or without the ".json" suffix.
func RemoveJSONSuffixFromParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := range c.Params {

			c.Params[i].Value = trimSuffixIfPresent(c.Params[i].Value, ".json")
			c.Params[i].Key = trimSuffixIfPresent(c.Params[i].Key, ".json")
		}

		c.Next()
	}
}
func trimSuffixIfPresent(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return strings.TrimSuffix(s, suffix)
	}
	return s
}
