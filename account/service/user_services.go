package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/memrizr/account/model"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository model.UserRepository
}


// USConfig will hold repositories that will eventually be injected into this
// this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService initializes a UserService with its repository layer dependencies
func NewUserService(userRepo model.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

// func NewUserService(c *USConfig) model.UserService {
// 	return &UserService{
// 		UserRepository: c.UserRepository,
// 	}
// }	


// Get retrieves a user based on their UUID
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	// if err != nil {
	// 	return nil, err
	// }
	return u, err
}



func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	panic("Method not implemeted")
}