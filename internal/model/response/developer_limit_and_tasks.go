package response

type DeveloperLimitAndTasks struct {
	MaxLimit    int   `json:"max_limit"`
	Fullness    int   `json:"fullness"`
	TaskIds     []int `json:"task_ids"`
	DeveloperId int   `json:"developer_id"`
}

type EfficiencyResponse struct {
	TotalWeekCount         int                      `json:"total_week_count"`
	DeveloperLimitAndTasks []DeveloperLimitAndTasks `json:"developer_limit_and_tasks"`
}
