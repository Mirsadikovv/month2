package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method is", r.Method)

	fmt.Fprintf(w, "<<<what?>>>")
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "<<< Welcome >>>")
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/index", indexHandler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
