// externalapi/fetch.go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type RandomOrgResponse struct {
	Data []int `json:"data"`
}

func fetchRandomOrgAPI(length string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	apiUrl := fmt.Sprintf("https://www.random.org/integers/?num=%s&min=1&max=100&col=1&base=10&format=plain&rnd=new", length)
	request, err := http.NewRequest("GET", apiUrl, nil)
	request = request.WithContext(ctx)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "json;charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data := strings.TrimSuffix(string(responseBody), "\n")
	responseSlice := strings.Split(data, "\n")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	return responseSlice, nil
}
