package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", GetTasksHandler)
	mux.HandleFunc("POST /tasks", CreateTaskHandler)

	log.Println("Server running on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
