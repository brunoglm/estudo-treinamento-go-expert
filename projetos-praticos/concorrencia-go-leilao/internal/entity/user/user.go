package user

import (
	"auction-go/internal/internalerror"
	"context"
)

type User struct {
	Id   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(
		ctx context.Context, userId string) (*User, *internalerror.InternalError)
}
