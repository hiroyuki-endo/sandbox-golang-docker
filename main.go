package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/microsoft/vscode-remote-try-go/controllers"
	"github.com/microsoft/vscode-remote-try-go/database"
	"github.com/microsoft/vscode-remote-try-go/repositories"
	"github.com/microsoft/vscode-remote-try-go/todo"
)

func main() {
	db := database.GormConnect()
	defer db.Close()
	db.AutoMigrate(&todo.Todo{})
	router := gin.Default()
	taskRepository := repositories.CreateTaskRepository(db)
	taskController := controllers.CreateTaskController(taskRepository, router)

	taskController.Endpoints()

	router.Run(":8080")
}
