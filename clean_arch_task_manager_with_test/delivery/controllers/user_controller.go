package controllers

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	hashedPassword, err := infrastructure.Generatepassword(user.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error hashing the password"})
		return
	}
	adduser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: user.Username,
		Password: string(hashedPassword),
		Role:     user.Role,
	}
	userid, err := uc.UserUsecase.RegisterUser(&adduser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error registering the user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userid": userid})

}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := uc.UserUsecase.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	c.Header("Authorization", jwtToken)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
