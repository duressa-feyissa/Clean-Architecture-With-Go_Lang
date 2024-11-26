package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

func Generatepassword(password string) (string,error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword),nil
}

func Checkpassword(userpassword, enteredpassword string)  error{
	err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(enteredpassword))
	return err
}