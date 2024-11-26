package repository_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo/mocks"
	"cleantaskmanager/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegisterUser(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionUser

	mockUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "User 1",
		Password: "Password 1",
		Role:     "user",
	}
	mockemptyUser := domain.User{}
	mockUserID := "12345"

	t.Run("success", func(t *testing.T) {

		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserID, nil).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUserRepository(databaseHelper, collectionName)

		err := ur.RegisterUser(&mockUser)

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockemptyUser, errors.New("Unexpected")).Once()

		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUserRepository(databaseHelper, collectionName)
		err := ur.RegisterUser(&mockUser)
		assert.Error(t, err)
		collectionHelper.AssertExpectations(t)
	})
}

func TestGetUserByID(t *testing.T) {
	// var databaseHelper *mocks.Database
	// var collectionHelper *mocks.Collection
	// var singleResultHelper *mocks.SingleResult
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	singleResultHelper := &mocks.SingleResult{}
	collectionName := domain.CollectionTask
	t.Run("success", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(singleResultHelper).Once()
		singleResultHelper.On("Decode", mock.AnythingOfType("*domain.User")).Return(nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewUserRepository(databaseHelper, collectionName)
		_, err := tr.GetUserByID(primitive.NewObjectID())
		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		singleResultHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(singleResultHelper).Once()
		singleResultHelper.On("Decode", mock.AnythingOfType("*domain.User")).Return(errors.New("Unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)
		tr := repository.NewUserRepository(databaseHelper, collectionName)
		_, err := tr.GetUserByID(primitive.NewObjectID())
		assert.Error(t, err)
		collectionHelper.AssertExpectations(t)
		singleResultHelper.AssertExpectations(t)
	})
}
