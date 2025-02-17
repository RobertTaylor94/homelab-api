package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"murvoth.co.uk/homeapi/models"
)

func HomeKitPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var homeStatus models.HomeStatus

		if err := json.Unmarshal(requestBody, &homeStatus); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		// for each dictionary of home data in homeStatus.Data
		// attempt to save the data to the database
		for _, d := range homeStatus.Data {
			if err := saveHomeStatus(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		c.JSON(200, "Saved")
	}
}

// saveTempStatus saves the timestamp and current temperature from a sensor to the temperature_sensors table
func saveHomeStatus(db *sql.DB, homeStatus *models.HomeStatusData) error {

	parsedTime, err := parseTimeHome(homeStatus.Timestamp)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	query := `INSERT INTO homekit.sensor_readings (room_id, temperature, humidity, recorded_at) VALUES ($1, $2, $3, $4) ON CONFLICT (room_id, recorded_at) DO NOTHING`
	_, err = db.Exec(query, 1, homeStatus.Temperature, homeStatus.Humidity, parsedTime)
	if err != nil {
		return fmt.Errorf("failed to save home status: %v", err)
	}

	return nil
}