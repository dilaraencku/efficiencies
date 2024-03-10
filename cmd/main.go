package main

import (
	"bufio"
	"efficientDevelopment/db"
	"efficientDevelopment/internal/client"
	"efficientDevelopment/internal/model"
	"efficientDevelopment/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go server.Init()

	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	if dbGorm == nil {
		panic("Database connection is nil")
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	providers := []model.Provider{
		{
			Id:       1,
			Name:     "First",
			Endpoint: "https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd",
		},
		{
			Id:       2,
			Name:     "Second",
			Endpoint: "https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143",
		},
	}

	inputChan := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter valid input: ")
		text, _ := reader.ReadString('\n')
		inputChan <- text
	}()

	userInput := <-inputChan

	if userInput == "getDatas\n" {
		client.ProcessCommand(providers)
	}

	sig := <-signalChannel

	log.Printf("%s exit signal \ndetected, server is closing <3", sig)
}
