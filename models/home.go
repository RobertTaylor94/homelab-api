package models

type HomeStatus struct {
	Data []HomeStatusData `json:"data"`
}

type HomeStatusData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
	SensorName  string  `json:"sensor_name"`
}
