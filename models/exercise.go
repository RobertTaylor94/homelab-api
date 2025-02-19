package models

// step count
type StepCount struct {
	Data StepCountData `json:"data"`
}

type StepCountData struct {
	Metrics []StepCountMetrics `json:"metrics"`
}

type StepCountMetrics struct {
	Name  string           `json:"name"`
	Units string           `json:"units"`
	Data  []StepCountDaily `json:"data"`
}

type StepCountDaily struct {
	Date     string  `json:"date"`
	Source   string  `json:"source"`
	Quantity float32 `json:"qty"`
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

// energy
type Energy struct {
	Data EnergyData `json:"data"`
}

type EnergyData struct {
	Metrics []EnergyMetrics `json:"metrics"`
}

type EnergyMetrics struct {
	Units string        `json:"units,omitempty"`
	Name  string        `json:"name,omitempty"`
	Data  []EnergyDaily `json:"data"`
}

type EnergyDaily struct {
	Date     string  `json:"date,omitempty"`
	Quantity float64 `json:"qty,omitempty"`
	Type     string
}
