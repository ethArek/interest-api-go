package dependencies

import (
	"github.com/ethArek/interest-api-go/app/services"
)

// Dependencies type
type Dependencies struct {
	InterestService *services.InterestService
}

// Deps instance of dependencies.Dependencies
var Deps *Dependencies

// CreateDependencies function
func CreateDependencies() {
	interestService := &services.InterestService{}

	Deps = &Dependencies{
		interestService,
	}
}
