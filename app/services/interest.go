package services

import (
	"github.com/ethArek/interest-api-go/app/types"
	"github.com/go-pg/pg"
)

// InterestService struct
type InterestService struct {
}

// GetInterests method
func (u *InterestService) GetInterests() []*types.Interest {
	return []*types.Interest{
		&types.Interest{
			ID:   "50fd10d3-3d38-451b-b75e-121030ce13de",
			Name: "interestname",
		},
	}
}

// GetInterestByID method
func (u *InterestService) GetInterestByID(id string) *types.Interest {
	return &types.Interest{
		ID:   "50fd10d3-3d38-451b-b75e-121030ce13de",
		Name: "interestname",
	}
}

// PostInterest method
func (u *InterestService) PostInterest(body types.CreateInterest) *types.Interest {
	db := pg.Connect(&pg.Options{
		Addr:     "interestapigo_postgres_1:5432",
		User:     "postgres",
		Password: "postgres_password",
	})

	defer db.Close()

	return &types.Interest{}
}
