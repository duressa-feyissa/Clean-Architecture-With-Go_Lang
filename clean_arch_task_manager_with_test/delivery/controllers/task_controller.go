package controllers

import (
	"cleantaskmanager/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	claims := &domain.Claims{
		UserID: userID,
		Role:   c.GetString("role"),
	}
	tasks, err := tc.TaskUsecase.GetTasks(c, claims)
	if err != nil || len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Message": "No tasks found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	claims := &domain.Claims{
		UserID: userID,
		Role:   c.GetString("role"),
	}
	id,_ := primitive.ObjectIDFromHex(c.Param("id"))
	task, err := tc.TaskUsecase.GetTask(c, claims, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) AddTask(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	claims := &domain.Claims{
		UserID: userID,
		Role:   c.GetString("role"),
	}
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adtask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
		UserID:      userID,
	}

	taskid,err := tc.TaskUsecase.AddTask(c, claims, &adtask)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error adding the task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully", "taskid": taskid})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	claims := &domain.Claims{
		UserID: userID,
		Role:   c.GetString("role"),
	}
	id,_ := primitive.ObjectIDFromHex(c.Param("id"))
	var task domain.UpdateTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.TaskUsecase.UpdateTask(c, claims, id, &task)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error updating the task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
	claims := &domain.Claims{
		UserID: userID,
		Role:   c.GetString("role"),
	}
	id,_ := primitive.ObjectIDFromHex(c.Param("id"))
	err := tc.TaskUsecase.DeleteTask(c, claims, id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error deleting the task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
