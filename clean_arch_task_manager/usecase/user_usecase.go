package usecase
import (
	"cleantaskmanager/domain"
	"cleantaskmanager/infrastructure"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (lu *loginUsecase) RegisterUser(user *domain.User) error {
	return lu.userRepository.RegisterUser(user)
}

func (lu *loginUsecase) LoginUser(user *domain.User) (string ,error) {
	jwttoken , error := infrastructure.GenerateToken(user)
	return jwttoken , error
	
}

func (lu *loginUsecase) GetUserByID(id string) (*domain.User, error) {
	return lu.userRepository.GetUserByID(id)
}

