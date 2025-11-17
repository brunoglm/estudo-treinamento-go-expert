package repository

import (
	"context"
	"fmt"
	"mongodb-lab/entity"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		coll: collection,
	}
}

func (s *UserRepository) Create(ctx context.Context, user *entity.User) (bson.ObjectID, error) {
	user.CreatedAt = time.Now()

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := s.coll.InsertOne(opCtx, user)
	if err != nil {
		return bson.NilObjectID, fmt.Errorf("Erro ao inserir usuário: %v", err)
	}

	id, _ := result.InsertedID.(bson.ObjectID)
	return id, nil
}

func (s *UserRepository) GetByID(ctx context.Context, id bson.ObjectID) (*entity.User, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user entity.User
	filter := bson.M{"_id": id}

	err := s.coll.FindOne(opCtx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("Usuário não encontrado")
	} else if err != nil {
		return nil, fmt.Errorf("Erro ao buscar usuário por ID: %v", err)
	}

	return &user, nil
}

func (s *UserRepository) UpdateName(ctx context.Context, id bson.ObjectID, newName string) (int64, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	// pode ser adicionado outros campos para update
	updateData := bson.M{
		"$set": bson.M{"name": newName},
	}

	result, err := s.coll.UpdateOne(opCtx, filter, updateData)
	if err != nil {
		return 0, fmt.Errorf("Erro ao atualizar usuário: %v", err)
	}

	return result.ModifiedCount, nil
}

func (s *UserRepository) Replace(ctx context.Context, id bson.ObjectID, updateUser *entity.User) (int64, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	updateUser.ID = id

	result, err := s.coll.ReplaceOne(opCtx, filter, updateUser)
	if err != mongo.ErrNoDocuments {
		return 0, fmt.Errorf("Usuário não encontrado para substituição")
	} else if err != nil {
		return 0, fmt.Errorf("Erro ao substituir usuário: %v", err)
	}

	return result.ModifiedCount, nil
}

func (s *UserRepository) Delete(ctx context.Context, id bson.ObjectID) (int64, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := s.coll.DeleteOne(opCtx, filter)
	if err != nil {
		return 0, fmt.Errorf("Erro ao deletar usuário: %v", err)
	}

	return result.DeletedCount, nil
}

func (s *UserRepository) UseDbTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	session, err := s.coll.Database().Client().StartSession()
	if err != nil {
		return fmt.Errorf("Erro ao iniciar sessão: %v", err)
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx context.Context) (interface{}, error) {
		err := fn(sessCtx)
		return nil, err
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return fmt.Errorf("Erro na transação: %v", err)
	}

	return nil
}
