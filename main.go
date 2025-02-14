package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// How to use:
// curl -X POST -H "Content-Type: application/json" -d '{"a": 10, "b": 5}' http://localhost:8080/arithmetic?op=add
// Send a POST request to /arithmetic?op=<operation> with a JSON body like:
// {
//    "a": 10,
//    "b": 5
// }
//
// Supported operations: add, subract, multiply, divide

type RequestInput struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type OperationResponse struct {
	Result float64 `json:"result"`
}

func arithmeticinputs(operation string, w http.ResponseWriter, r *http.Request) {
	var req RequestInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("Error decoding request %v", err)
		http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest)
		return
	}

	var result float64
	if operation == "add" {
		result = req.A + req.B
	} else if operation == "subtract" {
		result = req.A - req.B
	} else if operation == "multiply" {
		result = req.A * req.B
	} else if operation == "divide" {
		result = req.A / req.B
	} else {
		log.Printf("Operation doesnt exist")
		http.Error(w, "Operation doesnt exist", http.StatusBadRequest)
		return
	}

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

func main() {
	http.HandleFunc("/arithmetic", func(w http.ResponseWriter, r *http.Request) {
		op := r.URL.Query().Get("op") // Get the "op" query parameter
		arithmeticinputs(op, w, r)    // Call your handler function
	})
	port := os.Getenv("PORT") // Reads the value of the "PORT" environment variable
	if port == "" {           // Checks if the 'port' variable is empty
		port = "8080" // Default port if not specified
	}
	addr := ":" + port // Constructs the address string by prepending a colon to the port
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
