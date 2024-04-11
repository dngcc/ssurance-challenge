package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	slidingwindow "github.com/ssurance-challenge/internal/slidingwindow"
)

func Start(ctx context.Context, apiPort string, intervalInSeconds int, numberOfTicks int, saveFilePath string) error {
	windowDuration := time.Duration(intervalInSeconds) * time.Second
	slWindow := slidingwindow.NewCounter(ctx, windowDuration, uint(numberOfTicks), saveFilePath)

	http.HandleFunc("/counter", counter(slWindow))
	http.ListenAndServe(":"+apiPort, nil)

	return nil
}

func counter(sl *slidingwindow.Counter) func(resp http.ResponseWriter, _ *http.Request) {

	return func(resp http.ResponseWriter, _ *http.Request) {

		var counterResp = struct {
			Counter uint `json:"counter"`
		}{
			Counter: sl.IncreaseCount(),
		}

		b, err := json.Marshal(counterResp)
		if err != nil {
			http.Error(resp, "Failed to parse counter response", http.StatusInternalServerError)
			return
		}

		if _, err = resp.Write(b); err != nil {
			http.Error(resp, "Failed to write counter response", http.StatusInternalServerError)
		}
	}
}
