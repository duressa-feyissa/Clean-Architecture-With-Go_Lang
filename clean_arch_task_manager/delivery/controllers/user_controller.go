package controllers

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := infrastructure.Generatepassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	adduser := domain.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Role:     user.Role,
	}
	adduser = *domain.NewUser(adduser.Username, adduser.Password, adduser.Role)
	err = adduser.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = uc.UserUsecase.RegisterUser(&adduser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error registering the user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "userid": adduser.ID})

}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := uc.UserUsecase.GetUserByID(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = infrastructure.Checkpassword(result.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
		return
	}
	jwtToken, err := infrastructure.GenerateToken(result)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.Header("Authorization", jwtToken)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful","userid":result.ID})
}
