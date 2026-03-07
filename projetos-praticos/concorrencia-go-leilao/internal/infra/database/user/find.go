package userrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/user"
	internalerrors "auction-go/internal/errors"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user.User, *internalerrors.Error) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			message := fmt.Sprintf("User not found with this id = %s", userId)
			logger.Error(message, err)
			return nil, internalerrors.NewNotFoundError(message)
		}

		message := "Error trying to find user by userId"
		logger.Error(message, err)
		return nil, internalerrors.NewInternalServerError(message)
	}

	userEntity := &user.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
