package routers

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo"
	"github.com/gin-gonic/gin"
	"cleantaskmanager/repository"
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/usecase"
)

func NewUserRouter(db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := &controllers.UserController{
		UserUsecase: usecase.NewUserUsecase(ur),
	}
	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
}
