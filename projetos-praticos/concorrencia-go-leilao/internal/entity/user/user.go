package user

import (
	"auction-go/internal/errors"
	"context"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *errors.Error)
}
