package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	Text string `json:"text"`
}

var (
	msg  Message
	lock sync.Mutex
)

func main() {
	msg = Message{Text: "..."}

	http.HandleFunc("/", handleMessage)

	fmt.Println("Server started at :6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	switch r.Method {
	case http.MethodGet:
		lock.Lock()
		defer lock.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)

	case http.MethodPost:
		var newMsg Message
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &newMsg); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		lock.Lock()
		msg.Text = newMsg.Text
		lock.Unlock()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Message updated")

	case http.MethodHead:
		// Handle HEAD request
		lock.Lock()
		defer lock.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}