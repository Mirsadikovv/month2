package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	url := "http://localhost:8080/user"

	userBody := User{
		Id:    "58d5443f-f143-4988-82f3-0d02202723ff",
		Name:  "Tom",
		Email: "tommy13@mail.com",
		Phone: "998990019292",
	}

	marshalledBody, err := json.Marshal(userBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalledBody))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
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
