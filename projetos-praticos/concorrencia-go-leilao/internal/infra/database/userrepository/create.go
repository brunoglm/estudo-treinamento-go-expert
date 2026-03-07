package userrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/user"
	"auction-go/internal/errors"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, userEntity *user.User) *errors.Error {
	userEntityMongo := &UserEntityMongo{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}

	_, err := ur.Collection.InsertOne(ctx, userEntityMongo)
	if err != nil {
		message := "Error trying to insert user"
		logger.Error(message, err)
		return errors.NewInternalServerError(message)
	}

	return nil
}
