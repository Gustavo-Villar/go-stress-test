package entity

import "time"

type Report struct {
	TimeSpent          time.Duration
	RequestsMade       int
	SuccessfulRequests int
	FailedRequests     map[string]int
}

func NewReport(timeSpent time.Duration, requestsMade int, successfulRequests int, failedRequests map[string]int) *Report {
	return &Report{
		TimeSpent:          timeSpent,
		RequestsMade:       requestsMade,
		SuccessfulRequests: successfulRequests,
		FailedRequests:     failedRequests,
	}
}
