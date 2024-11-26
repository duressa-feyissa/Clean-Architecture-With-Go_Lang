package usecase

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (lu *UserUsecase) RegisterUser(user *domain.User) (primitive.ObjectID, error) {
	return user.ID, lu.userRepository.RegisterUser(user)
}

func (lu *UserUsecase) LoginUser(user *domain.Login) (string, error) {
	error := infrastructure.Checkpassword(user.Password, user.Password)
	if error != nil {
		return "", error
	}
	login,err := lu.userRepository.GetUserByID(user.ID)
	if err != nil {
		return "", err
	}
	jwttoken, err := infrastructure.GenerateToken(login)
	return jwttoken, err
}

func (lu *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	return lu.userRepository.GetUserByID(id)
}
