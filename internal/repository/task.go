package repository

import (
	"efficientDevelopment/internal/model"
	"gorm.io/gorm"
	"log"
)

func TaskBulkInsert(db *gorm.DB, task []model.Task) {

	result := db.CreateInBatches(task, len(task))
	if result.Error != nil {
		log.Fatalf("Error happened when inserting tasks")
	}
}

func TaskList(db *gorm.DB) (task []model.Task) {

	result := db.Order("total_time DESC").Find(&task)

	if result.Error != nil {
		log.Printf("Error happend getting developer efficiencies ")
	}

	return task
}
