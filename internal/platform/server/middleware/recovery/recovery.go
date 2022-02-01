package recovery

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Middleware returns a middleware that recovers from any panics and writes a 500 if there was one.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[Recovery] panic recovered: %v\n", err)
				c.Abort()
				c.AbortWithStatus(500)
			}
		}()

		// Process next middleware (or request handler)
		c.Next()
	}
}
