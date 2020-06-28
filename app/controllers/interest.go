package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethArek/interest-api-go/app/dependencies"
	"github.com/ethArek/interest-api-go/app/types"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// GetInterests controller
func GetInterests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// interests := dependencies.Deps.InterestService.GetInterests()
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres_password@interestapigo_postgres_1:5432/postgres")
	if err != nil {
		panic("sth")
	}

	var interest *types.Interest
	errorek := db.QueryRow(context.Background(), "SELECT id, name, language FROM interests WHERE name = $1", "interest").Scan(&interest)
	if errorek != nil {
		fmt.Println(errorek)
		json.NewEncoder(w).Encode(errorek)
	} else {
		json.NewEncoder(w).Encode(interest)

	}

}

// GetInterest controller
func GetInterest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	interestID := vars["id"]

	interest := dependencies.Deps.InterestService.GetInterestByID(interestID)
	json.NewEncoder(w).Encode(interest)
}

func PostInsert(w http.ResponseWriter, r *http.Request) {

}
