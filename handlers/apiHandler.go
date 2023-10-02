package handlers

import (
	"encoding/json"
	"fmt"
	"httpserver/data"
	"httpserver/externalapi"
	"net/http"
	"strconv"
	"sync"
)

type ResponseData struct {
	Data  []string `json:"data"`
	Stdev float64  `json:"stdev"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request")

	requests := r.URL.Query().Get("requests")
	length := r.URL.Query().Get("length")

	response, _ := doRequests(w, requests, length)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func doRequests(w http.ResponseWriter, r, l string) ([]ResponseData, error) {
	numberOfRequests, err := strconv.Atoi(r)
	if err != nil {
		return []ResponseData{}, err
	}

	var wg sync.WaitGroup

	resultsCh := make(chan ResponseData, numberOfRequests)

	for i := 0; i < numberOfRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			randomApiData, err := externalapi.FetchRandomOrgAPI(l)
			if err != nil {
				fmt.Printf("Error fetching data: %v\n", err)
				return
			}

			// Calculate the standard deviation
			stdDev := data.CalculateSTDEV(randomApiData)

			response := ResponseData{
				Stdev: stdDev,
				Data:  randomApiData,
			}

			resultsCh <- response
		}()
	}

	wg.Wait()

	// Close the results channel
	close(resultsCh)

	var responses []ResponseData
	for response := range resultsCh {
		responses = append(responses, response)
	}

	return responses, nil
}
