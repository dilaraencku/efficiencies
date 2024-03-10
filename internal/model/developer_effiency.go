package model

import "time"

type DeveloperEfficiency struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	DeveloperName string    `gorm:"unique" json:"developer_name"`
	Hour          int       `json:"hour"`
	Difficulty    int       `json:"difficulty"`
	TotalValue    int       `json:"total_value"`
	CreatedAt     time.Time `json:"created_at"`
}
