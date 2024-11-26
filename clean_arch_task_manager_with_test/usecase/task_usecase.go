package usecase

import (
	"cleantaskmanager/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	taskRepository domain.TaskRepository
}
func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
	}
}

func (tu *TaskUsecase) GetTasks(c context.Context, claims *domain.Claims) ([]domain.Task, error) {
	return tu.taskRepository.GetTasks(c, claims)
}

func (tu *TaskUsecase) GetTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) (*domain.Task, error) {
	return tu.taskRepository.GetTask(c, claims, id)
}

func (tu *TaskUsecase) AddTask(c context.Context, claims *domain.Claims ,task *domain.Task) (primitive.ObjectID, error) {
	return task.ID,tu.taskRepository.AddTask(c,claims, task)
}

func (tu *TaskUsecase) UpdateTask(c context.Context, claims *domain.Claims, id primitive.ObjectID, task *domain.UpdateTask) error {
	return tu.taskRepository.UpdateTask(c, claims, id, task)
}

func (tu *TaskUsecase) DeleteTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) error {
	return tu.taskRepository.DeleteTask(c, claims, id)
}

