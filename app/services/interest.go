package services

import (
	"github.com/ethArek/interest-api-go/app/types"
)

// InterestService struct
type InterestService struct {
}

// GetInterests method
func (u *InterestService) GetInterests() []*types.Interest {
	return []*types.Interest{
		&types.Interest{
			ID:       1,
			Name:     "interestname",
			Language: 1,
		},
	}
}

// GetInterestByID method
func (u *InterestService) GetInterestByID(id string) *types.Interest {
	return &types.Interest{
		ID:       1,
		Name:     "interestname",
		Language: 1,
	}
}

// PostInterest method
func (u *InterestService) PostInterest(body types.CreateInterest) *types.Interest {
	return &types.Interest{}
}
