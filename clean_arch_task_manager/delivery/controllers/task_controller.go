package controllers

import (
	"cleantaskmanager/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	claims := &domain.Claims{
		UserID: c.GetString("user_id"),
		Role:   c.GetString("role"),
	}
	tasks, err := tc.TaskUsecase.GetTasks(c, claims)
	if err != nil || len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	claims := &domain.Claims{
		UserID: c.GetString("user_id"),
		Role:   c.GetString("role"),
	}
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTask(c, claims, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) AddTask(c *gin.Context) {
	claims := &domain.Claims{
		UserID: c.GetString("user_id"),
		Role:   c.GetString("role"),
	}
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	duedate, er := time.Parse(time.RFC3339, task.DueDate)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	adtask := domain.Task{
		Title:       task.Title,
		Description: task.Description,
		DueDate:     duedate.Format(time.RFC3339),
		Status:      task.Status,
		UserID:      claims.UserID,
	}

	adtask = *domain.NewTask(claims.UserID, task.Title, task.Description, duedate.Format(time.RFC3339), task.Status)
	err := tc.TaskUsecase.AddTask(c, claims, &adtask)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error adding the task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully", "taskid": adtask.ID})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	claims := &domain.Claims{
		UserID: c.GetString("user_id"),
		Role:   c.GetString("role"),
	}
	id := c.Param("id")
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
	claims := &domain.Claims{
		UserID: c.GetString("user_id"),
		Role:   c.GetString("role"),
	}
	id := c.Param("id")
	err := tc.TaskUsecase.DeleteTask(c, claims, id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error deleting the task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
