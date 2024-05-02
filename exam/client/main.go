package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}
type TodoList struct {
	TaskId       string
	TaskName     string
	TaskTime     time.Time
	TaskComment  string
	TaskPriority string
	UserId       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {

	url_create := "http://localhost:8080/create_task"
	url_update := "http://localhost:8080/update_task"
	url_delete := "http://localhost:8080/delete_task"
	url_select := "http://localhost:8080/select_task"
	url_priority := "http://localhost:8080/select_priority"
	url_register := "http://localhost:8080/register"
	url_login := "http://localhost:8080/login"

	currentTime := time.Now()

	userBody := User{
		Id:    "110ec58a-a0f2-4ac4-8393-c866d813b8d1",
		Name:  "Tom",
		Age:   31,
		Phone: "998903334413",
	}

	TaskBody := TodoList{
		TaskId:       "110ec58a-a0f2-4ac4-8393-c866d813b8d1",
		TaskName:     "obed",
		TaskTime:     currentTime.Add(24 * time.Hour),
		TaskComment:  "ne zabud",
		TaskPriority: "high",
		UserId:       "d1b7857d-94d0-4107-b419-f55f86ad9775",
	}
	///////////////////////////////// login login

	marshalledBody, err := json.Marshal(userBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req, err := http.NewRequest("POST", url_login, bytes.NewBuffer(marshalledBody))
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

	///////////////////// registr registr

	marshalledBody2, err := json.Marshal(userBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req2, err := http.NewRequest("POST", url_register, bytes.NewBuffer(marshalledBody2))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp2.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body2))

	//////////////////// create create

	marshalledBody3, err := json.Marshal(TaskBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req3, err := http.NewRequest("POST", url_create, bytes.NewBuffer(marshalledBody3))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp3, err := client.Do(req3)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp3.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body3))

	////////////////////update update

	marshalledBody4, err := json.Marshal(TaskBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req4, err := http.NewRequest("PUT", url_update, bytes.NewBuffer(marshalledBody4))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp4, err := client.Do(req4)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp4.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body4, err := io.ReadAll(resp4.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body4))

	//////////////////// select select

	marshalledBody5, err := json.Marshal(TaskBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req5, err := http.NewRequest("GET", url_select, bytes.NewBuffer(marshalledBody5))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp5, err := client.Do(req5)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp5.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body5, err := io.ReadAll(resp5.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body5))

	//////////////////// priority select

	marshalledBody6, err := json.Marshal(TaskBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req6, err := http.NewRequest("GET", url_priority, bytes.NewBuffer(marshalledBody6))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp6, err := client.Do(req6)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp6.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body6, err := io.ReadAll(resp6.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body6))

	//////////////////// delete delete

	marshalledBody7, err := json.Marshal(TaskBody)
	if err != nil {
		log.Fatal("impossible to marshal user", err)
	}

	req7, err := http.NewRequest("DELETE", url_delete, bytes.NewBuffer(marshalledBody7))
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp7, err := client.Do(req7)
	if err != nil {
		fmt.Println("Error getting response", err)
		return
	}

	defer resp7.Body.Close()
	fmt.Println("code :", resp.StatusCode)

	body7, err := io.ReadAll(resp7.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}
	fmt.Println("Response : ", string(body7))

	////////////////////
}
