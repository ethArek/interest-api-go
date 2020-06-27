package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ethArek/interest-api-go/app/dependencies"
	"github.com/gorilla/mux"
)

// GetInterests controller
func GetInterests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	interests := dependencies.Deps.InterestService.GetInterests()
	json.NewEncoder(w).Encode(interests)
}

// GetInterest controller
func GetInterest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	interestID := vars["id"]

	interest := dependencies.Deps.InterestService.GetInterestByID(interestID)
	json.NewEncoder(w).Encode(interest)
}
