package middleware

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

var lastRequest = make(map[string]time.Time)

func RateLimitMiddleware(limit time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        now := time.Now()

        if last, exists := lastRequest[ip]; exists && now.Sub(last) < limit {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            return
        }

        lastRequest[ip] = now
        c.Next()
    }
}
