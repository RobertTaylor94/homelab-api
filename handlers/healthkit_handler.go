package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
			if err := saveHeartData(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}

func HealthKitStepsPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()

		var stepData models.StepCount

		if err := json.Unmarshal(requestBody, &stepData); err != nil {
			fmt.Printf("failed to unmarshal json: %v\n", err)
		}

		for _, d := range stepData.Data.Metrics[0].Data {
			if err := saveStepData(db, &d); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}
