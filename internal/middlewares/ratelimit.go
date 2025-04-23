package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func RateLimitMiddleware(redisClient *redis.Client, rateLimit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, err := strconv.Atoi(rateLimit)
		if err != nil {
			limit = 100 // default
		}

		key := "rate_limit:" + c.ClientIP()

		count, err := redisClient.Incr(c, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		if count == 1 {
			redisClient.Expire(c, key, time.Minute)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(429, gin.H{
				"error":      "Too many requests",
				"rate_limit": limit,
				"remaining":  0,
			})
			return
		}

		c.Header("X-RateLimit-Limit", strconv.Itoa(limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(limit-int(count)))
		c.Next()
	}
}
