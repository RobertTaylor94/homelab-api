package handlers

import (
	"database/sql"
	"fmt"
	"math"

	"murvoth.co.uk/homeapi/models"
)

// saveHeartData is the helper function that executes a SQL query to save heart rate data to the database. Will ignore previously written data if any entries overlap with existing database data.
func saveHeartData(db *sql.DB, d *models.HeartRates) error {
	// parse time recieved in json to the correct format for the database
	parsedTime, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// sql query to insert heart rate data into the database, if entry exists with the same timestamp ignore the request to insert data.
	query := `INSERT INTO healthkit.heart_data (recorded_at, avg_heart_rate) VALUES ($1, $2) ON CONFLICT (recorded_at) DO NOTHING`
	_, err = db.Exec(query, parsedTime, math.Round(float64(d.Average)))
	if err != nil {
		return fmt.Errorf("failed to save heart data: %v", err)
	}

	return nil
}

// saveStepData is the helper function that executes a SQL query to save daily step count data to the database. Will update previously stored data if an updated entry is recieved.
func saveStepData(db *sql.DB, d *models.StepCountDaily) error {
	// parse time recieved in json to the correct format for the database
	parsedTime, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// sql query to insert step data into the database, if entry exists with the same timestamp, update existing record with new data
	query := `INSERT INTO healthkit.step_data (recorded_at, quantity) VALUES ($1, $2) ON CONFLICT (recorded_at) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err = db.Exec(query, parsedTime, math.Round(float64(d.Quantity)))
	if err != nil {
		return fmt.Errorf("failed to save step data: %v", err)
	}

	return nil
}

// saveEnergyData is the helper function that executes a SQL query to save energy data. Will update previously stored data if an updated entry is recieved.
func saveEnergyData(db *sql.DB, d *models.EnergyDaily) error {
	// parse time recieved in json to the correct format for the database
	parsedTime, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// sql query to insert step data into the database, if entry exists with the same timestamp, update existing record with new data
	query := `INSERT INTO healthkit.energy_data (recorded_at, quantity, energy_type) VALUES ($1, $2, $3) ON CONFLICT (recorded_at, energy_type) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err = db.Exec(query, parsedTime, math.Round(float64(d.Quantity)), d.Type)
	if err != nil {
		return fmt.Errorf("failed to save step data: %v", err)
	}

	return nil
}

// saveSleepData is the helper function that executes a SQL query to save sleep data. Will update previously stored data if an updated entry is recieved.
func saveSleepData(db *sql.DB, d *models.SleepAnalysisDaily) error {
	// parse time recieved in json to the correct format for the database
	recorded_at, err := parseTimeHealth(d.Date)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %v", err)
	}

	// sql query to insert step data into the database, if entry exists with the same timestamp, update existing record with new data
	query := `
	INSERT INTO healthkit.sleep_data
  		(recorded_at, in_bed_start, in_bed_end, rem, core, asleep, awake, deep, sleep_start, sleep_end)
	VALUES
  		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (recorded_at) DO UPDATE
	SET
  		in_bed_start = EXCLUDED.in_bed_start,
  		in_bed_end   = EXCLUDED.in_bed_end,
  		rem          = EXCLUDED.rem,
  		core         = EXCLUDED.core,
  		asleep       = EXCLUDED.asleep,
  		awake        = EXCLUDED.awake,
  		deep         = EXCLUDED.deep,
  		sleep_start  = EXCLUDED.sleep_start,
  		sleep_end    = EXCLUDED.sleep_end;
	`
	_, err = db.Exec(query, recorded_at, d.InBedStart, d.InBedEnd, d.REM, d.Core, d.Asleep, d.Awake, d.Deep, d.SleepStart, d.SleepEnd)
	if err != nil {
		return fmt.Errorf("failed to save step data: %v", err)
	}

	return nil
}
