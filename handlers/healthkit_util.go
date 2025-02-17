package handlers

import (
	"database/sql"
	"fmt"
	"math"

	"murvoth.co.uk/homeapi/models"
)

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

func saveStepData(db *sql.DB, d *models.StepCountDaily) error {
	parsedTime, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	query := `INSERT INTO healthkit.step_data (recorded_at, quantity) VALUES ($1, $2) ON CONFLICT (recorded_at) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err = db.Exec(query, parsedTime, math.Round(float64(d.Quantity)))
	if err != nil {
		return fmt.Errorf("failed to save step data: %v", err)
	}

	return nil
}
