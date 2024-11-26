package routers
import (
	"cleantaskmanager/mongo"
	"github.com/gin-gonic/gin"
	"cleantaskmanager/repository"
	"cleantaskmanager/domain"
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/usecase"
)
func NewTaskRouter(db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controllers.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr),
	}
	group.GET("/tasks", tc.GetTasks)
	group.GET("/tasks/:id", tc.GetTask)
	group.PUT("/tasks/:id", tc.UpdateTask)
	group.POST("/tasks", tc.AddTask)
	group.DELETE("/tasks/:id", tc.DeleteTask)
}