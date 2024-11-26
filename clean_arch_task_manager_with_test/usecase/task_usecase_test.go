package usecase_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"cleantaskmanager/usecase"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseSuite struct {
	suite.Suite
	mockRepo    *mocks.TaskRepository
	taskUseCase *usecase.TaskUsecase
}

func (suite *TaskUseCaseSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepository)
	suite.taskUseCase = usecase.NewTaskUsecase(suite.mockRepo).(*usecase.TaskUsecase)
}

func (suite *TaskUseCaseSuite) TestAddTask() {
	tasks := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
		UserID:      primitive.NewObjectID(),
	}
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	suite.mockRepo.On("AddTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("*domain.Task")).Return(nil)
	_,err := suite.taskUseCase.AddTask(context.Background(), &claims, &tasks)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetTask() {
	taskID := primitive.NewObjectID()
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	suite.mockRepo.On("GetTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(&domain.Task{}, nil)
	task, err := suite.taskUseCase.GetTask(context.Background(), &claims, taskID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), task)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetTasks() {
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	suite.mockRepo.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return([]domain.Task{}, nil)
	tasks, err := suite.taskUseCase.GetTasks(context.Background(), &claims)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), tasks)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestUpdateTask() {
	taskID := primitive.NewObjectID()
	task := domain.UpdateTask{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     "2021-01-01T00:00:00Z",
		Status:      "pending",
	}
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	suite.mockRepo.On("UpdateTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID"), mock.AnythingOfType("*domain.UpdateTask")).Return(nil)
	err := suite.taskUseCase.UpdateTask(context.Background(), &claims, taskID, &task)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID()
	claims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}
	suite.mockRepo.On("DeleteTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)
	err := suite.taskUseCase.DeleteTask(context.Background(), &claims, taskID)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}