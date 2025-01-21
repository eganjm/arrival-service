package handlers

import (
	"encoding/json"
	"net/http"

	"in-flight-service/models"
)

// GetArrivingFlights retrieves arriving flight information
func GetInFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := models.FetchInFlights()
	if err != nil {
		http.Error(w, "Unable to fetch arriving flights", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
