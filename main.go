package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/microsoft/vscode-remote-try-go/database"
	"github.com/microsoft/vscode-remote-try-go/todo"
)

func main() {
	db := database.GormConnect()
	defer db.Close()
	db.AutoMigrate(&todo.Todo{})

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		todos := []todo.Todo{}
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		var newTodo todo.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&newTodo)
		c.JSON(http.StatusOK, newTodo.ID)
	})

	router.DELETE("/todos", func(c *gin.Context) {
		todos := []todo.Todo{}
		db.Find(&todos)
		db.Delete(&todos)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		todo := todo.Todo{}
		db.Where("id = ?", id).Find(&todo)
		db.Delete(&todo)
	})

	router.Run(":8080")
}
