package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type ResponseData struct {
	Data   []string `json:"data"`
	Stddev float64  `json:"stddev"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request")

	requests := r.URL.Query().Get("requests")
	length := r.URL.Query().Get("length")

	response, err := doRequests(requests, length)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error during the requestes %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func doRequests(r, l string) ([]ResponseData, error) {
	numberOfRequests, err := strconv.Atoi(r)
	if err != nil {
		return []ResponseData{}, err
	}

	var wg sync.WaitGroup

	resultsCh := make(chan ResponseData, numberOfRequests)
	errChan := make(chan error, numberOfRequests)

	for i := 0; i < numberOfRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			randomApiData, err := fetchRandomOrgAPI(l)
			if err != nil {
				fmt.Printf("Error fetching data: %v\n", err)
				errChan <- err
				return
			}

			stdDev, err := calculateSTDDEV(randomApiData)
			if err != nil {
				fmt.Printf("Error caulculating stddev: %v\n", err)
				errChan <- err
				return
			}

			response := ResponseData{
				Stddev: stdDev,
				Data:   randomApiData,
			}

			resultsCh <- response
		}()
	}

	wg.Wait()
	close(resultsCh)
	close(errChan)

	if len(errChan) != 0 {
		return nil, errors.New("error during the data fetching")
	}

	return calculateStdev(resultsCh)
}

func calculateStdev(resultsCh chan ResponseData) ([]ResponseData, error) {
	var allSets []string

	var responses []ResponseData
	for response := range resultsCh {
		responses = append(responses, response)
		allSets = append(allSets, response.Data...)
	}

	stddevOfAllSets, err := calculateSTDDEV(allSets)
	if err != nil {
		fmt.Printf("Error caulculating stddev: %v\n", err)
		return []ResponseData{}, err
	}

	return append(responses, ResponseData{
		Data:   allSets,
		Stddev: stddevOfAllSets,
	}), nil
}
