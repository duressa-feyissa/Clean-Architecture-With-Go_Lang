package repository

import (
	"cleantaskmanager/domain"
	"cleantaskmanager/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func(tr *taskRepository) AddTask(c context.Context, claims *domain.Claims, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	return err	

}

func(tr *taskRepository) GetTasks(c context.Context, claims *domain.Claims) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	filter := bson.D{{Key: "userid", Value: claims.UserID}}
	if claims.Role == "admin" {
		filter = bson.D{}
	}
	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	var tasks []domain.Task
	for cursor.Next(c) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, err
}

func(tr *taskRepository) GetTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) (*domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: claims.UserID}}
	if claims.Role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	var result domain.Task
	err := collection.FindOne(c, filter).Decode(&result)
	return &result, err
}

func(tr *taskRepository) UpdateTask(c context.Context, claims *domain.Claims, id primitive.ObjectID, task *domain.UpdateTask) error {
	collection := tr.database.Collection(tr.collection)
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: claims.UserID}}
	if claims.Role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: task.Title},
			{Key: "description", Value: task.Description},
			{Key: "duedate", Value: task.DueDate},
			{Key: "status", Value: task.Status},
		}},
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func(tr *taskRepository) DeleteTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) error {
	collection := tr.database.Collection(tr.collection)
	filter := bson.D{{Key: "id", Value: id}, {Key: "userid", Value: claims.UserID}}
	if claims.Role == "admin" {
		filter = bson.D{{Key: "id", Value: id}}
	}
	_, err := collection.DeleteOne(c, filter)
	return err
}
