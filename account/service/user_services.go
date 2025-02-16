package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/moheddine-belhaj/Translator/account/model"
	"github.com/moheddine-belhaj/Translator/account/model/apperrors"
)


type UserService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected into this
// this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

// Signup reaches our to a UserRepository to sign up the user.
// UserRepository Create should handle checking for user exists conflicts
func (s *UserService) Signup(ctx context.Context, u *model.User) error {
    pw, err := hashPassword(u.Password)

    if err != nil {
        log.Printf("Unable to signup user for email: %v\n", u.Email)
        return apperrors.NewInternal()
    }

    // now I realize why I originally used Signup(ctx, email, password)
    // then created a user. It's somewhat un-natural to mutate the user here
    u.Password = pw
    if err := s.UserRepository.Create(ctx, u); err != nil {
        return err
    }

    // ...

    return nil
}
