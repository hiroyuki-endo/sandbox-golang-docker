package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microsoft/vscode-remote-try-go/models"
	"github.com/microsoft/vscode-remote-try-go/repositories"
)

type TaskController struct {
	taskRepository *repositories.TaskRepository
	router         *gin.Engine
}

func CreateTaskController(taskRepository *repositories.TaskRepository, router *gin.Engine) *TaskController {
	taskController := new(TaskController)
	taskController.taskRepository = taskRepository
	taskController.router = router
	return taskController
}

func (tc *TaskController) Endpoints() {
	tc.getTasks()
	tc.createTask()
	tc.deleteTasks()
	tc.deleteTaskById()
	tc.doneTask()
}

func (tc *TaskController) getTasks() {
	tc.router.GET("/todos", func(c *gin.Context) {
		todos := tc.taskRepository.All()
		c.JSON(http.StatusOK, todos)
	})
}

func (tc *TaskController) createTask() {
	tc.router.POST("/todos", func(c *gin.Context) {
		var newTodo models.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tc.taskRepository.Create(&newTodo)
		c.JSON(http.StatusOK, newTodo.ID)
	})
}

func (tc *TaskController) doneTask() {
	tc.router.POST("/todos/:id/done", func(c *gin.Context) {
		id := c.Param("id")
		numId, _ := strconv.Atoi(id)
		doneTask := tc.taskRepository.FindById(numId)
		doneTask.Done()
		tc.taskRepository.Save(doneTask)
	})
}

func (tc *TaskController) deleteTasks() {
	tc.router.DELETE("/todos", func(c *gin.Context) {
		tc.taskRepository.DeleteAll()
	})
}

func (tc *TaskController) deleteTaskById() {
	tc.router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		numId, _ := strconv.Atoi(id)
		tc.taskRepository.DeleteById(numId)
	})
}
