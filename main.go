package main

import (
	"encoding/json" // Package to handle JSON encoding and decoding
	"fmt"           // Package for formatted I/O, like printing to the console
	"log"           // Package for logging errors and other information
	"net/http"      // Package for building HTTP servers and clients
)

// OperationRequest defines the structure for the request body
type OperationRequest struct {
	A float64 `json:"a"` // Field 'A' of type float64, mapped to JSON key "a"
	B float64 `json:"b"` // Field 'B' of type float64, mapped to JSON key "b"
}

// OperationResponse defines the structure for the response body
type OperationResponse struct {
	Result float64 `json:"result"` // Field 'Result' of type float64, mapped to JSON key "result"
}

// addHandler is the handler function for the "/add" endpoint
func addHandler(w http.ResponseWriter, r *http.Request) {
	var req OperationRequest // Declare a variable 'req' of type OperationRequest

	//  TODO: Handle potential errors during decoding more gracefully. Consider logging the error.
	// Example:
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	log.Printf("Error decoding request: %v", err) // Log the error
	// 	http.Error(w, "Bad Request - Invalid JSON", http.StatusBadRequest) // Send a more user-friendly error
	// 	return
	// }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // Decode the JSON from the request body into the 'req' variable
		http.Error(w, err.Error(), http.StatusBadRequest) // If there's an error during decoding, return a HTTP 400 error with the error message
		return                                            // Exit the handler function
	}

	result := req.A + req.B                   // Calculate the sum of 'A' and 'B' from the request
	resp := OperationResponse{Result: result} // Create a response object with the calculated result

	w.Header().Set("Content-Type", "application/json") // Set the Content-Type header to "application/json"
	//  TODO: Handle potential errors during encoding.
	// Example:
	// enc := json.NewEncoder(w)
	// if err := enc.Encode(resp); err != nil {
	// 	log.Printf("Error encoding response: %v", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	json.NewEncoder(w).Encode(resp) // Encode the response object to JSON and write it to the response writer
}

// main is the entry point of the program
func main() {
	http.HandleFunc("/add", addHandler) // Register the 'addHandler' function to handle requests to the "/add" endpoint

	//  TODO: Consider using environment variables for the port number.
	// Example:
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080" // Default port if not specified
	// }
	// addr := ":" + port
	// fmt.Printf("Server listening on port %s\n", port)
	// log.Fatal(http.ListenAndServe(addr, nil))
	fmt.Println("Server listening on port 8080") // Print a message to the console
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the HTTP server and listen on port 8080, log.Fatal will exit if there is an error
}
