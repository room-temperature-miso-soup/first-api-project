package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RequestInput struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type OperationResponse struct {
	Result float64 `json:"result"`
}

func addinputs(w http.ResponseWriter, r *http.Request) {
	var req RequestInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Error decoding request %v", err)
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}
	result := req.A + req.B
	fmt.Printf("result: %v\n", result)
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func subtractinputs(w http.ResponseWriter, r *http.Request) {
	var req RequestInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Error decodiong request %v", err)
		http.Error(w, "Bad request - Invalid JSON", http.StatusBadRequest)
		return
	}
	result := req.A - req.B
	resp := OperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("result: %v\n", result)
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/subtract", subtractinputs)
	http.HandleFunc("/add", addinputs)
	port := os.Getenv("PORT") // Reads the value of the "PORT" environment variable
	if port == "" {           // Checks if the 'port' variable is empty
		port = "8080" // Default port if not specified
	}
	addr := ":" + port // Constructs the address string by prepending a colon to the port
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
