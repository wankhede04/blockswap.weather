package weatherservice

import (
	"net/http"
	"time"

	"github.com/wankhede04/blockswap.weather/weather-srv/db"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware restricts user to call API within configured time frame
func (s *WeatherService) RateLimitMiddleware() gin.HandlerFunc {
	// Use a semaphore to limit concurrent requests
	return func(c *gin.Context) {
		s.semaphore.Add(1)
		defer s.semaphore.Done()

		membership, ok := c.Get("membership")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		m, ok := membership.(db.Membership)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		currentTime := time.Now().Unix()

		if currentTime-m.LastCall < 12 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		timeDifference := currentTime - m.LastCall
		if timeDifference%12 > 2 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Window passed"})
			c.Abort()
			return
		}
		c.Next()
	}
}
