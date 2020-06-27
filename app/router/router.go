package router

import (
	"github.com/ethArek/interest-api-go/app/controllers"
	"github.com/gorilla/mux"
)

// GetRouter - get default router
func GetRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/interests", controllers.GetInterests).Methods("GET")
	r.HandleFunc("/interests/{id}", controllers.GetInterest).Methods("GET")

	return r
}
