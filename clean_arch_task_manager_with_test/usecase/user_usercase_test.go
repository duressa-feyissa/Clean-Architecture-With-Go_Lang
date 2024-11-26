package usecase_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"cleantaskmanager/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
	usecase  *usecase.UserUsecase
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(mocks.UserRepository)
	suite.usecase = usecase.NewUserUsecase(suite.mockRepo).(*usecase.UserUsecase)
}
func (suite *UserUsecaseSuite) TearDownSuite() {
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestGetUserByID() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.mockRepo.On("GetUserByID", user.ID).Return(&user, nil)
	result, err := suite.usecase.GetUserByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
}

func (suite *UserUsecaseSuite) TestLoginUser() {
	user:= domain.Login{ ID: primitive.NewObjectID(), Password: "123"}
	suite.mockRepo.On("LoginUser",mock.AnythingOfType("*domain.Login")).Return("jtw_token",nil)
	jwt,err := suite.usecase.LoginUser(&user)
	suite.Nil(err)
	suite.NotEmpty(jwt)

}

func (suite *UserUsecaseSuite) TestRegisterUser() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.mockRepo.On("RegisterUser", &user).Return(nil)
	_,err := suite.usecase.RegisterUser(&user)
	assert.NoError(suite.T(), err)
	
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
