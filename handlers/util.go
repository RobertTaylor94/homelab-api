package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "OK")
	}
}

// helper to parse timestamp from string sent in JSON
func parseTimeHome(timestamp string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing time: %v", err)
	}
	return parsedTime, nil
}

// helper to parse timestamp from string sent in JSON
func parseTimeHealth(timestamp string) (time.Time, error) {
	layout := "2006-01-02 15:04:05 -0700"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return time.Now(), fmt.Errorf("error parsing time: %v", err)
	}
	return parsedTime, nil
}
