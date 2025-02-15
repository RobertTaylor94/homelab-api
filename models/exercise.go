package models

// step count
type StepCount struct {
	Data StepCountData `json:"data"`
}

type StepCountData struct {
	Metrics []StepCountMetrics `json:"metris"`
}

type StepCountMetrics struct {
	Name  string           `json:"name"`
	Units string           `json:"units"`
	Data  []StepCountDaily `json:"data"`
}

type StepCountDaily struct {
	Date     string `json:"date"`
	Source   string `json:"source"`
	Quantity int    `json:"qty"`
}

// exercise minutes
type ExerciseTime struct {
	Data ExerciseTimeData `json:"data"`
}

type ExerciseTimeData struct {
	Metrics []ExerciseTimeMetrics `json:"metrics"`
}

type ExerciseTimeMetrics struct {
	Name  string              `json:"name"`
	Units string              `json:"units"`
	Data  []ExerciseTimeDaily `json:"data"`
}

type ExerciseTimeDaily struct {
	Quantity int    `json:"qty"`
	Date     string `json:"date"`
	Source   string `json:"source"`
}

// active energy
type ActiveEnergy struct {
	Data ActiveEnergyData `json:"data"`
}

type ActiveEnergyData struct {
	Metrics ActiveEnergyMetrics `json:"metrics"`
}

type ActiveEnergyMetrics struct {
	Units string              `json:"units,omitempty"`
	Name  string              `json:"name,omitempty"`
	Data  []ActiveEnergyDaily `json:"data"`
}

type ActiveEnergyDaily struct {
	Date     string  `json:"date,omitempty"`
	Quantity float64 `json:"qty,omitempty"`
	Source   string  `json:"source,omitempty"`
}
