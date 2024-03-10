package client

import (
	"efficientDevelopment/db"
	"efficientDevelopment/internal/model"
	"efficientDevelopment/internal/model/provider_models"
	"efficientDevelopment/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type ResponseProcessor interface {
	ProcessResponse(response *http.Response) (interface{}, error)
}

type ProviderFirst struct {
	ResponseData provider_models.FirstProvider
}

type ProviderFirstResponseProcessor struct{}

func (p ProviderFirstResponseProcessor) ProcessResponse(response *http.Response) (myInterface interface{}, err error) {
	var firstProviderResponse []provider_models.FirstProvider

	if err := json.NewDecoder(response.Body).Decode(&firstProviderResponse); err != nil {
		log.Printf("decoding error %v", err)
	}

	var task []model.Task
	for _, e := range firstProviderResponse {
		task = append(task, model.Task{
			ExternalId:   e.Id,
			Duration:     e.Sure,
			Difficulty:   e.Zorluk,
			TotalTime:    e.Sure * e.Zorluk,
			CreatedAt:    time.Now(),
			ProviderName: "first",
		})
	}

	repository.TaskBulkInsert(db.MyDB, task)

	return firstProviderResponse, err
}

type ProviderSecondResponseProcessor struct{}

func (p ProviderSecondResponseProcessor) ProcessResponse(response *http.Response) (myInterface interface{}, err error) {
	var secondProviderResponse []provider_models.SecondProvider

	if err := json.NewDecoder(response.Body).Decode(&secondProviderResponse); err != nil {
		log.Printf("decoding error %v", err)
	}
	var task []model.Task
	for _, e := range secondProviderResponse {
		task = append(task, model.Task{
			ExternalId:   e.Id,
			Duration:     e.EstimatedDuration,
			Difficulty:   e.Value,
			TotalTime:    e.EstimatedDuration * e.Value,
			CreatedAt:    time.Now(),
			ProviderName: "second",
		})
	}

	repository.TaskBulkInsert(db.MyDB, task)

	return secondProviderResponse, err
}

func ProcessCommand(providers []model.Provider) {

	var wg sync.WaitGroup

	results := make(chan string, len(providers))

	for _, e := range providers {
		wg.Add(1)
		go func(provider model.Provider) {
			defer wg.Done()
			result := GetProvidersData(provider)
			results <- result
		}(e)
	}
	wg.Wait()

	close(results)
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func GetProvidersData(provider model.Provider) string {

	response, err := http.Get(provider.Endpoint)
	if err != nil {
		fmt.Println("Error :", err)
	}

	processor := GetResponseProcessor(provider.Name)
	if processor != nil {
		model, err := processor.ProcessResponse(response)
		if err != nil {
			fmt.Println("Error :", model, err)
		}
	}

	return fmt.Sprintf("Process done for provider %s", provider.Name)
}

func GetResponseProcessor(providerName string) ResponseProcessor {
	switch providerName {
	case "First":
		return ProviderFirstResponseProcessor{}
	case "Second":
		return ProviderSecondResponseProcessor{}
	default:
		return nil
	}
}
