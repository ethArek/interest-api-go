package app

import "github.com/ethArek/interest-api-go/app/services"

type AppContainer struct {
	userService *services.UserService
}

func GetAppContainer() *AppContainer {

}
