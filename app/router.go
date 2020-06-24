package router

import (
	"github.com/ethArek/interest-api-go/app/handlers"
	"github.com/gorilla/mux"
)

// GetRouter - get default router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)

	return r
}
