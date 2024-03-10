package repository

import (
	"efficientDevelopment/internal/model"
	"gorm.io/gorm"
	"log"
)

func DeveloperEfficiencyList(db *gorm.DB) (developerEff []model.DeveloperEfficiency) {

	result := db.Order("total_value DESC").Find(&developerEff)

	if result.Error != nil {
		log.Printf("Error happend getting developer efficiencies ")
	}

	return developerEff
}
