package models

type Workouts struct {
	Data WorkoutsData `json:"data"`
}

type WorkoutsData struct {
	Workouts []Workout `json:"workouts"`
}

type Workout struct {
	Name          string          `json:"name"`
	HeartRateData []HeartRates    `json:"heartRateData"`
	Distance      WorkoutDistance `json:"distance,omitempty"`
	ActiveEnergy  WorkoutEnergy   `json:"activeEnergyBurned"`
	Start         string          `json:"start"`
	End           string          `json:"end"`
}

type WorkoutDistance struct {
	Units    string  `json:"units"`
	Quantity float64 `json:"qty"`
}

type WorkoutEnergy struct {
	Quantity float64 `json:"qty"`
	Units    string  `json:"units"`
}
