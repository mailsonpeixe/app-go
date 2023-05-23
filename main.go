package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting the app.\n")
	//Handlers to readiness and liveness
	http.HandleFunc("/healthz/ready", readinessHandler)
	http.HandleFunc("/healthz/live", livenessHandler)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the app is ready to serve requests
	// and return a 200 OK response if so
	// If not, return a 500 Internal Server Error response
	if isReady() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	// Return a 200 OK response to indicate that the app is running
	w.WriteHeader(http.StatusOK)
}

func isReady() bool {
	// Implement your own readiness check here
	// For example, you might check if the app has successfully connected to a database
	return true
}
