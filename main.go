package main

import (
	"net/http"

	"github.com/ethArek/interest-api-go/app/dependencies"

	"github.com/ethArek/interest-api-go/app/router"
)

func main() {
	dependencies.CreateDependencies()
	router := router.GetRouter()

	http.ListenAndServe(":8080", router)
}
