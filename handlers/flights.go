package handlers

import (
	"encoding/json"
	"net/http"

	"arrivals-service/models"
)

// GetArrivingFlights retrieves arriving flight information
func GetArrivingFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := models.FetchArrivingFlights()
	if err != nil {
		http.Error(w, "Unable to fetch arriving flights", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
