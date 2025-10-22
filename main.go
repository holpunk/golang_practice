package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello, Cloud Native World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
