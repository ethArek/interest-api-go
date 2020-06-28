package main

import (
	"fmt"
	"net/http"

	"github.com/ethArek/interest-api-go/app/dependencies"

	"github.com/ethArek/interest-api-go/app/router"
)

func main() {
	dependencies.CreateDependencies()
	router := router.GetRouter()

	http.ListenAndServe(":8080", router)
	fmt.Println("listening on 8080")
}
