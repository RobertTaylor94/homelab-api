package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"murvoth.co.uk/homeapi/models"
)

// HealthKitHeartPost is the handler for saving heart rate data to the database. Will ignore already saved results if a batch request is made.
func HealthKitHeartPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var heartData models.HeartRate

		// unmarshal raw request data in heartData model
		if err := json.Unmarshal(requestBody, &heartData); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		// for each entry in heartData, execute a sql query to save the data to the database
		for _, d := range heartData.Data.Metrics[0].Data {
			if err := saveHeartData(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}

// HealthKitStepsPost is the handler for saving daily step count rate data to the database. Will update existing values if sent multiple times.
func HealthKitStepsPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var stepData models.StepCount

		// unmarshal raw request data unto stepData model
		if err := json.Unmarshal(requestBody, &stepData); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		// for each entry in stepData, execute a sql query to save data to the database
		for _, d := range stepData.Data.Metrics[0].Data {
			if err := saveStepData(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}

func HealthKitEnergyPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var energyData models.Energy

		// unmarshal raw request data into energyData model
		if err := json.Unmarshal(requestBody, &energyData); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		for _, e := range energyData.Data.Metrics {
			energyType := e.Name
			for _, d := range e.Data {
				d.Type = energyType
				if err := saveEnergyData(db, &d); err != nil {
					fmt.Println(err.Error())
					c.JSON(http.StatusInternalServerError, err.Error())
				}
			}
		}
	}
}
