package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/ssurance-challenge/internal/server"
)

const (
	defaultPort     = "8080"
	defaultFilePath = "../sliding_window_counter.gob"
)

func main() {

	ctx := context.Background()

	var numberOfTicks int
	var err error
	var intervalInSeconds int

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = defaultPort
	}

	if os.Getenv("NUMBER_OF_TICKS") != "" {
		numberOfTicks, err = strconv.Atoi(os.Getenv("NUMBER_OF_TICKS"))
		if err != nil {
			log.Printf("invalid number of ticks %s", os.Getenv("NUMBER_OF_TICKS"))
			os.Exit(0)
		}

	}

	if os.Getenv("INTERVAL_IN_SECONDS") != "" {
		intervalInSeconds, err = strconv.Atoi(os.Getenv("INTERVAL_IN_SECONDS"))
		if err != nil {
			log.Printf("invalid number of seconds %s", os.Getenv("INTERVAL_IN_SECONDS"))
			os.Exit(0)
		}

	}

	saveFilePath := os.Getenv("SAVE_FILE_PATH")
	if saveFilePath == "" {
		saveFilePath = defaultFilePath
	}

	server.Start(ctx, apiPort, numberOfTicks, intervalInSeconds, saveFilePath)
}
