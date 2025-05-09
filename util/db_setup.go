package util

import (
	"database/sql"
	"fmt"
)

func InitHealthKit(db *sql.DB) error {
	stmts := []string{
		`CREATE SCHEMA IF NOT EXISTS healthkit`,
		`CREATE TABLE IF NOT EXISTS healthkit.heart_data (
            recorded_at TIMESTAMPTZ NOT NULL,
            avg_heart_rate INTEGER NOT NULL,
            PRIMARY KEY (recorded_at)
        )`,
		`CREATE TABLE IF NOT EXISTS healthkit.step_data (
            recorded_at TIMESTAMPTZ NOT NULL,
            quantity INTEGER NOT NULL,
            PRIMARY KEY (recorded_at)
        )`,
		`CREATE TABLE IF NOT EXISTS healthkit.energy_data (
            recorded_at TIMESTAMPTZ NOT NULL,
            quantity INTEGER NOT NULL,
            energy_type TEXT NOT NULL,
            PRIMARY KEY (recorded_at, energy_type)
        )`,
		`CREATE TABLE IF NOT EXISTS healthkit.sleep_data (
            recorded_at TIMESTAMPTZ NOT NULL,
            in_bed_start TIMESTAMPTZ,
            in_bed_end TIMESTAMPTZ,
            rem DOUBLE PRECISION,
            core DOUBLE PRECISION,
            asleep DOUBLE PRECISION,
            awake DOUBLE PRECISION,
            deep DOUBLE PRECISION,
            sleep_start TIMESTAMPTZ,
            sleep_end TIMESTAMPTZ,
            PRIMARY KEY (recorded_at)
        )`,
	}
	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("InitHealthKit: failed to execute statement %q: %v", stmt, err)
		}
	}
	return nil
}

func InitHomeKit(db *sql.DB) error {
	stmts := []string{
		`CREATE SCHEMA IF NOT EXISTS homekit`,
		`CREATE TABLE IF NOT EXISTS homekit.sensor_readings (
            room_id INTEGER NOT NULL,
            temperature DOUBLE PRECISION,
            humidity DOUBLE PRECISION,
            recorded_at TIMESTAMPTZ NOT NULL,
            PRIMARY KEY (room_id, recorded_at)
        )`,
	}
	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("InitHomeKit: failed to execute statement %q: %v", stmt, err)
		}
	}
	return nil
}
