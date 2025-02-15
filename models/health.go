package models

type Health struct {
	StepData  StepCount
	HeartData HeartRate
	SleepData SleepAnalysis
	Workouts  Workouts
}
