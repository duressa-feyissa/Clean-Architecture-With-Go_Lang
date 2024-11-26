package repository_test

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo/mocks"
	"cleantaskmanager/repository"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositorySuite struct {
	suite.Suite
	databaseHelper     *mocks.Database
	collectionHelper   *mocks.Collection
	cursorHelper       *mocks.Cursor
	singleResultHelper *mocks.SingleResult
}

func (suite *TaskRepositorySuite) SetupTest() {
	suite.databaseHelper = &mocks.Database{}
	suite.collectionHelper = &mocks.Collection{}
	suite.cursorHelper = &mocks.Cursor{}
	suite.singleResultHelper = &mocks.SingleResult{}
}
func (suite *TaskRepositorySuite) TearDownSuite() {
	suite.collectionHelper.AssertExpectations(suite.T())
	suite.databaseHelper.AssertExpectations(suite.T())
	suite.cursorHelper.AssertExpectations(suite.T())
	suite.singleResultHelper.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestAddTask_Success() {
	collectionName := domain.CollectionTask

	mockTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
		UserID:      primitive.NewObjectID(),
	}
	mockClaims := domain.Claims{}

	suite.collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Task")).Return("12345", nil).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	ur := repository.NewTaskRepository(suite.databaseHelper, collectionName)
	err := ur.AddTask(context.Background(), &mockClaims, &mockTask)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TestAddTask_Error() {
	collectionName := domain.CollectionTask

	mockemptyTask := domain.Task{}
	mockClaims := domain.Claims{}

	suite.collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(mockemptyTask, errors.New("Unexpected")).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	ur := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	err := ur.AddTask(context.Background(), &mockClaims, &mockemptyTask)

	assert.Error(suite.T(), err)

	suite.collectionHelper.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestGetTasks_Success() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}

	suite.collectionHelper.On("Find", mock.Anything, mock.Anything).Return(suite.cursorHelper, nil).Once()
	suite.cursorHelper.On("Close", mock.Anything).Return(nil).Once()
	suite.cursorHelper.On("Next", mock.Anything).Return(false).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	_, err := tr.GetTasks(context.Background(), &mockClaims)
	suite.NoError(err)
}

func (suite *TaskRepositorySuite) TestGetTasks_Error() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}

	suite.collectionHelper.On("Find", mock.Anything, mock.Anything).Return(suite.cursorHelper, errors.New("Unexpected")).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	_, err := tr.GetTasks(context.Background(), &mockClaims)

	suite.Error(err)
}

func (suite *TaskRepositorySuite) TestGetTask_Success() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}

	suite.collectionHelper.On("FindOne", mock.Anything, mock.AnythingOfType("primitive.D")).Return(suite.singleResultHelper).Once()
	suite.singleResultHelper.On("Decode", mock.AnythingOfType("*domain.Task")).Return(nil).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	_, err := tr.GetTask(context.Background(), &mockClaims, primitive.NewObjectID())

	suite.NoError(err)


}

func (suite *TaskRepositorySuite) TestGetTask_Error() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{UserID: primitive.NewObjectID(), Role: "user"}

	suite.collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(suite.singleResultHelper).Once()
	suite.singleResultHelper.On("Decode", mock.AnythingOfType("*domain.Task")).Return(errors.New("Unexpected")).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	_, err := tr.GetTask(context.Background(), &mockClaims, primitive.NewObjectID())

	suite.Error(err)


}

func (suite *TaskRepositorySuite) TestUpdateTask_Success() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	mockTask := domain.UpdateTask{Title: "Updated Task"}
	updateresult := &mongo.UpdateResult{}

	suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, nil).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	err := tr.UpdateTask(context.Background(), &mockClaims, primitive.NewObjectID(), &mockTask)

	suite.NoError(err)

	
}

func (suite *TaskRepositorySuite) TestUpdateTask_Error() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	mockTask := domain.UpdateTask{Title: "Updated Task"}
	updateresult := &mongo.UpdateResult{}

	suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateresult, errors.New("Unexpected")).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	err := tr.UpdateTask(context.Background(), &mockClaims, primitive.NewObjectID(), &mockTask)

	suite.Error(err)

	
}

func (suite *TaskRepositorySuite) TestDeleteTask_Success() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	var id int64 = 1

	suite.collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(id, nil).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	err := tr.DeleteTask(context.Background(), &mockClaims, primitive.NewObjectID())

	suite.NoError(err)

	
}

func (suite *TaskRepositorySuite) TestDeleteTask_Error() {
	collectionName := domain.CollectionTask
	mockClaims := domain.Claims{}
	var id int64 = 1

	suite.collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(id, errors.New("Unexpected")).Once()
	suite.databaseHelper.On("Collection", collectionName).Return(suite.collectionHelper)

	tr := repository.NewTaskRepository(suite.databaseHelper, collectionName)

	err := tr.DeleteTask(context.Background(), &mockClaims, primitive.NewObjectID())

	suite.Error(err)
	
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
