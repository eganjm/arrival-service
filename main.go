package main

import (
	"log"
	"net/http"

	"in-flight-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Add CORS middleware
	r.Use(corsMiddleware)

	// Define routes
	r.HandleFunc("/in-flight", handlers.GetInFlights).Methods("GET")

	log.Println("Starting server on :8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
