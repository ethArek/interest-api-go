package services

import "github.com/ethArek/interest-api-go/app/types"

// UserService struct
type UserService struct {
}

// GetUsers method
func (u *UserService) GetUsers() []*types.User {
	return []*types.User{
		&types.User{
			ID:   "50fd10d3-3d38-451b-b75e-121030ce13de",
			Name: "username",
		},
	}
}

// GetUserByID method
func (u *UserService) GetUserByID(id string) *types.User {
	return &types.User{
		ID:   "50fd10d3-3d38-451b-b75e-121030ce13de",
		Name: "username",
	}
}
