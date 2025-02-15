package models

type HeartRate struct {
	Data HeartRateData `json:"data"`
}

type HeartRateData struct {
	Metrics []HeartRateMetrics `json:"metrics"`
}

type HeartRateMetrics struct {
	Units string       `json:"units"`
	Data  []HeartRates `json:"data"`
}

type HeartRates struct {
	Source  string `json:"source"`
	Average float32    `json:"Avg"`
	Date    string `json:"date"`
}
