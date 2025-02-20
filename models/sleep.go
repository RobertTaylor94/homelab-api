package models

type SleepAnalysis struct {
	Data SleepAnalysisData `json:"data"`
}

type SleepAnalysisData struct {
	Metrics []SleepAnalysisMetrics `json:"metrics"`
}

type SleepAnalysisMetrics struct {
	Name  string               `json:"name"`
	Units string               `json:"units"`
	Data  []SleepAnalysisDaily `json:"data"`
}

type SleepAnalysisDaily struct {
	Date       string  `json:"date"`
	InBedStart string  `json:"inBedStart"`
	InBedEnd   string  `json:"inBedEnd"`
	SleepStart string  `json:"sleepStart"`
	SleepEnd   string  `json:"sleepEnd"`
	InBed      float64 `json:"inBed"`
	Asleep     float64 `json:"asleep"`
	Awake      float64 `json:"awake"`
	Core       float64 `json:"core"`
	Deep       float64 `json:"deep"`
	REM        float64 `json:"rem"`
}
