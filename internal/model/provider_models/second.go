package provider_models

type SecondProvider struct {
	Id                string `json:"id"`
	Value             int    `json:"value"`
	EstimatedDuration int    `json:"estimated_duration"`
}
