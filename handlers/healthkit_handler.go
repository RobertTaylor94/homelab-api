package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"murvoth.co.uk/homeapi/models"
)

func HealthKitHeartPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var heartData models.HeartRate

		if err := json.Unmarshal(requestBody, &heartData); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		for _, d := range heartData.Data.Metrics[0].Data {
			fmt.Println("Attempting to save health data...")
			if err := saveHeartData(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}

func saveHeartData(db *sql.DB, d *models.HeartRates) error {
	parsedTime, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	query := `INSERT INTO healthkit.heart_data (recorded_at, avg_heart_rate) VALUES ($1, $2) ON CONFLICT (recorded_at) DO NOTHING`
	_, err = db.Exec(query, parsedTime, math.Round(float64(d.Average)))
	if err != nil {
		return fmt.Errorf("failed to save heart data: %v", err)
	}

	return nil
}
