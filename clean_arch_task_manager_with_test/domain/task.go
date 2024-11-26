package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "Tasks"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
}
type UpdateTask struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	DueDate     string `validate:"required"`
	Status      string `validate:"required"`
}

type TaskRepository interface {
	AddTask(c context.Context, claims *Claims, task *Task) error
	GetTasks(c context.Context, claims *Claims) ([]Task, error)
	GetTask(c context.Context, claims *Claims, id primitive.ObjectID) (*Task, error)
	UpdateTask(c context.Context, claims *Claims, id primitive.ObjectID, task *UpdateTask) error
	DeleteTask(c context.Context, claims *Claims, id primitive.ObjectID) error
}

type TaskUsecase interface {
	AddTask(c context.Context, claims *Claims, task *Task) (primitive.ObjectID, error)
	GetTasks(c context.Context, claims *Claims) ([]Task, error)
	GetTask(c context.Context, claims *Claims, id primitive.ObjectID) (*Task, error)
	UpdateTask(c context.Context, claims *Claims, id primitive.ObjectID, task *UpdateTask) error
	DeleteTask(c context.Context, claims *Claims, id primitive.ObjectID) error
}
