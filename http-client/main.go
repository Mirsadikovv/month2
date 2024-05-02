// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	url := "http://localhost:8080/user"

// 	req, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request", err)
// 		return
// 	}

// 	client := &http.Client{
// 		Transport: nil,
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 		},
// 		Jar:     nil,
// 		Timeout: 0,
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error getting response", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println("code :", resp.StatusCode)

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body", err)
// 		return
// 	}

// 	fmt.Println("Response : ", string(body))

// }

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/tv"
	requestData := []byte("example")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}

	fmt.Println("Response : ", string(body))

}
