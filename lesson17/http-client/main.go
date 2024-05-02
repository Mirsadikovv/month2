package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/index"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error while creating request")

	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error while response", err)
	}
	defer resp.Body.Close()
	fmt.Println("code: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while readin response body", err)
		return
	}
	fmt.Println("response: ", string(body))

}
