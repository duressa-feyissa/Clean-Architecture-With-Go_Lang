package repository

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) RegisterUser(user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func (ur *UserRepository) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}
