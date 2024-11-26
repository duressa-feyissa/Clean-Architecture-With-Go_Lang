package controllers_test

import (
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	taskController *controllers.TaskController
	mockUsecase    *mocks.TaskUsecase
	router         *gin.Engine
}

func (suite *TaskControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(mocks.TaskUsecase)
	suite.taskController = &controllers.TaskController{TaskUsecase: suite.mockUsecase}
	suite.router = gin.Default()
	suite.router.GET("/tasks", suite.taskController.GetTasks)
	suite.router.POST("/tasks", suite.taskController.AddTask)
	suite.router.PUT("/tasks/:id", suite.taskController.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.taskController.DeleteTask)
}

func (suite *TaskControllerSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestGetTasks() {

	suite.Run("List of Tasks", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		tasks := []domain.Task{
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 1",
				Description: "Task Description",
				DueDate:     time.Now(),
				Status:      "Pending",
			},
			{
				ID:          primitive.NewObjectID(),
				Title:       "Task 2",
				Description: "Task Description",
				DueDate:     time.Now(),
				Status:      "Pending",
			},
		}
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", nil)
		suite.mockUsecase.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return(tasks, nil).Once()
		suite.taskController.GetTasks(ctx)
		suite.Equal(200, w.Code)
	})

	suite.Run("Tasks not found", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		suite.mockUsecase.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return([]domain.Task{}, errors.New("errors")).Once()
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", nil)
		suite.taskController.GetTasks(ctx)
		expected, err := json.Marshal(gin.H{"Message": "No tasks found"})
		suite.Nil(err)

		suite.Equal(string(expected), w.Body.String())

		suite.Equal(404, w.Code)
	})

}

func (suite *TaskControllerSuite) TestGetTask() {

	suite.Run("Task found", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		task := domain.Task{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Task Description",
			DueDate:     time.Now(),
			Status:      "Pending",
		}
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", nil)
		suite.mockUsecase.On("GetTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(&task, nil).Once()
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: task.ID.Hex()})
		suite.taskController.GetTask(ctx)
		suite.Equal(200, w.Code)
	})

	suite.Run("Task not found", func() {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/register", nil)
		suite.mockUsecase.On("GetTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(&domain.Task{}, errors.New("error")).Once()
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "123456"})
		suite.taskController.GetTask(ctx)
		expected, err := json.Marshal(gin.H{"Message": "Task not found"})
		suite.Nil(err)

		suite.Equal(string(expected), w.Body.String())

		suite.Equal(404, w.Code)
	})
}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
