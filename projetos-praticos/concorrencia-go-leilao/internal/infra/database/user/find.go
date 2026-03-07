package userrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/user"
	"auction-go/internal/internalerror"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user.User, *internalerror.InternalError) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			message := fmt.Sprintf("User not found with this id = %s", userId)
			logger.Error(message, err)
			return nil, internalerror.NewNotFoundError(message)
		}

		message := "Error trying to find user by userId"
		logger.Error(message, err)
		return nil, internalerror.NewInternalServerError(message)
	}

	userEntity := &user.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
