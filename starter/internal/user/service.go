package user

import (
	"context"

	"github.com/jinzhu/copier"
	"starter/internal/ent"
)

type Service struct {
	db *ent.Client
}

func NewService(db *ent.Client) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateUser(ctx context.Context) (*User, error) {
	var newUser User

	rawUser, err := s.db.User.
		Create().
		SetName("name").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&newUser, &rawUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil

}
