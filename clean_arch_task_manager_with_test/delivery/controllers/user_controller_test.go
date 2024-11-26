package controllers_test

import (
	"bytes"
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	router         *gin.Engine
	userUsecase    *mocks.UserUsecase
	userController *controllers.UserController
}

func (suite *UserControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.userUsecase = new(mocks.UserUsecase)
	suite.userController = &controllers.UserController{UserUsecase: suite.userUsecase}
	suite.router = gin.Default()
	suite.router.POST("/register", suite.userController.Register)
	suite.router.POST("/login", suite.userController.Login)
}

func (suite *UserControllerSuite) TearDownTest() {
	suite.userUsecase.AssertExpectations(suite.T())

}

func (suite *UserControllerSuite) TestRegister() {
	suite.Run("Registration_success", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := domain.User{}
		userid := primitive.NewObjectID()
		suite.userUsecase.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(userid, nil).Once()
		body, _ := json.Marshal(user)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		suite.userController.Register(ctx)
		expected, err := json.Marshal(gin.H{"message": "User registered successfully", "userid": userid})
		suite.Nil(err)
		suite.Equal(201, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	suite.Run("Registration_Error", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		user := domain.User{}
		suite.userUsecase.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(primitive.ObjectID{}, errors.New("error")).Once()
		body, _ := json.Marshal(user)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		suite.userController.Register(ctx)
		expected, err := json.Marshal(gin.H{"message": "Error registering the user"})
		suite.Nil(err)
		suite.Equal(http.StatusNotAcceptable, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
	suite.Run("Request_Error", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("POST", "/register", nil)
		suite.userController.Register(ctx)
		expected, err := json.Marshal(gin.H{"error": "Invalid input"})
		suite.Nil(err)

		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})

}

func (suite *UserControllerSuite) TestLogin() {
	suite.Run("Login_Success", func() {
		user := domain.Login{ID: primitive.NewObjectID(), Password: "password"}
		jwtToken := "some.jwt.token"
		suite.userUsecase.On("LoginUser", mock.AnythingOfType("*domain.Login")).Return(jwtToken, nil).Once()
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusOK, w.Code)
	})
	suite.Run("Login_Error", func() {
		user := domain.Login{ID: primitive.NewObjectID(), Password: "password"}
		suite.userUsecase.On("LoginUser", mock.AnythingOfType("*domain.Login")).Return("", errors.New("error")).Once()
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)
		suite.Equal(http.StatusUnauthorized, w.Code)
	})
	suite.Run("Request_Error", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("POST", "/login", nil)
		suite.userController.Login(ctx)
		expected, err := json.Marshal(gin.H{"error": "EOF"})
		suite.Nil(err)
		suite.Equal(400, w.Code)
		suite.Equal(string(expected), w.Body.String())
	})
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
