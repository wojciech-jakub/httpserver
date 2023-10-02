// externalapi/fetch.go
package externalapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RandomOrgResponse struct {
	Data []int `json:"data"`
}

func FetchRandomOrgAPI(length string) ([]string, error) {
	apiUrl := fmt.Sprintf("https://www.random.org/integers/?num=%s&min=1&max=100&col=1&base=10&format=plain&rnd=new", length)
	fmt.Println(apiUrl)
	request, err := http.NewRequestWithContext(context.Background(), "GET", apiUrl, nil)

	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "json;charset=UTF-8")
	request.Header.Add("Accept", "*/*")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(responseBody))
	fmt.Println()
	//fmt.Printf("%#v", sResponseBody)
	fmt.Println()

	data := strings.TrimSuffix(string(responseBody), "\n")
	responseSlice := strings.Split(data, "\n")

	fmt.Printf("%#v", responseSlice)
	//fmt.Println()
	//fmt.Printf("%#v", responseBody)

	defer response.Body.Close()

	return responseSlice, nil
}
