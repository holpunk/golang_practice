package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to represent the data you want to return in the API response
type ResponseData struct {
	Message string `json:"message"`
}

// Define a handler function for your API endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Create an instance of your response data
	data := ResponseData{
		Message: "Hello from the Go API!",
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the response
	w.Write(jsonData)
}

func main() {
	// Define the API endpoint and handler function
	http.HandleFunc("/api", apiHandler)

	// Start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
