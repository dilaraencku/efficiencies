package model

import "time"

type Task struct {
	Id           int       `json:"id"`
	ExternalId   string    `json:"external_id"`
	Duration     int       `json:"hour"`
	Difficulty   int       `json:"difficulty"`
	TotalTime    int       `json:"total_time"`
	CreatedAt    time.Time `json:"created_at"`
	ProviderName string    `json:"provider_name"`
}
