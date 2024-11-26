package infrastructure

import (
	"cleantaskmanager/domain"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *domain.User) (string, error) {

	// Generate JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Minute * 40).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		print("there is error ")
	}

	return jwtToken, err

}
