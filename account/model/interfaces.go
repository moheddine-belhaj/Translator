package model

import (
	"context"

	"github.com/google/uuid"
)


type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

// with in regards to producing JWTs as string
type TokenService interface {
    NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
    Create(ctx context.Context, u *User) error
}
