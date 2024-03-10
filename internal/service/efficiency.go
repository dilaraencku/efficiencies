package service

import (
	"efficientDevelopment/db"
	"efficientDevelopment/internal/model/response"
	"efficientDevelopment/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetData(c echo.Context) error {

	efResp, totalWeek := GetDeveloperEfficiency()

	resp := response.EfficiencyResponse{
		DeveloperLimitAndTasks: efResp,
		TotalWeekCount:         totalWeek,
	}

	return c.JSON(http.StatusOK, resp)
}

func GetDeveloperEfficiency() (developerLimit []response.DeveloperLimitAndTasks, totalWeek int) {

	developerEfficiency := repository.DeveloperEfficiencyList(db.MyDB)
	tasks := repository.TaskList(db.MyDB)

	developersRealLimit := make(map[int]int)

	for _, e := range developerEfficiency {
		developerLimit = append(developerLimit, response.DeveloperLimitAndTasks{
			MaxLimit:    e.TotalValue,
			DeveloperId: e.Id,
		})

		developersRealLimit[e.Id] = e.TotalValue
	}

	weekCount := 1

	for taskIndex, task := range tasks {

		developerLimitCounter := 0

		for index := range developerLimit {

			if developerLimit[index].MaxLimit > developerLimit[index].Fullness+tasks[taskIndex].TotalTime {
				developerLimit[index].Fullness = developerLimit[index].Fullness + tasks[taskIndex].TotalTime
				developerLimit[index].TaskIds = append(developerLimit[index].TaskIds, tasks[taskIndex].Id)
				break
			} else {
				developerLimitCounter++
			}

		}

		if developerLimitCounter == len(developerLimit) {
			weekCount++
			for index := range developerLimit {
				developerLimit[index].Fullness = 0
			}

			tasks = append(tasks, task)
		}

	}

	return developerLimit, weekCount
}
