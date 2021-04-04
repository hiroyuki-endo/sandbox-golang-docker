package controllers

import (
	"fmt"
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
	tc.getTask()
	tc.createTask()
	tc.deleteTasks()
	tc.deleteTaskById()
	tc.startTask()
	tc.doneTask()
}

func (tc *TaskController) getTasks() {
	tc.router.GET("/todos", func(c *gin.Context) {
		todos := tc.taskRepository.All()
		c.JSON(http.StatusOK, todos)
	})
}

func (tc *TaskController) getTask() {
	tc.router.GET("/todos/:id", func(c *gin.Context) {
		task, err1 := tc.taskByPathId(c)
		if err1 != nil {
			fmt.Println(err1)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	})
}

func (tc *TaskController) createTask() {
	tc.router.POST("/todos", func(c *gin.Context) {
		var newTodo models.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTodo.Status = "Ready"
		tc.taskRepository.Create(&newTodo)
		c.JSON(http.StatusOK, newTodo.ID)
	})
}

func (tc *TaskController) startTask() {
	tc.router.POST("/todos/:id/start", func(c *gin.Context) {
		startTask, err1 := tc.taskByPathId(c)
		if err1 != nil {
			fmt.Println(err1)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
			return
		}
		err2 := startTask.Start()
		if err2 != nil {
			fmt.Println(err2)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
		tc.taskRepository.Save(startTask)
	})
}

func (tc *TaskController) doneTask() {
	tc.router.POST("/todos/:id/done", func(c *gin.Context) {
		doneTask, err1 := tc.taskByPathId(c)
		if err1 != nil {
			fmt.Println(err1)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
			return
		}
		err2 := doneTask.Done()
		if err2 != nil {
			fmt.Println(err2)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}
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

func (tc *TaskController) taskByPathId(c *gin.Context) (*models.Todo, error) {
	id := c.Param("id")
	numId, _ := strconv.Atoi(id)
	todo, err := tc.taskRepository.FindById(numId)
	if err != nil {
		return todo, err
	}
	return todo, nil
}
