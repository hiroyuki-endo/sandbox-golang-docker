package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/microsoft/vscode-remote-try-go/todo"
)

type TaskRepository struct {
	db *gorm.DB
}

func CreateTaskRepository(db *gorm.DB) *TaskRepository {
	newTaskRepository := new(TaskRepository)
	newTaskRepository.db = db
	return newTaskRepository
}

func (tr *TaskRepository) All() *[]todo.Todo {
	todos := []todo.Todo{}
	tr.db.Find(&todos)
	return &todos
}

func (tr *TaskRepository) Create(newTodo *todo.Todo) int {
	tr.db.Create(newTodo)
	return newTodo.ID
}

func (tr *TaskRepository) DeleteAll() {
	todos := []todo.Todo{}
	tr.db.Find(&todos)
	tr.db.Delete(&todos)
}

func (tr *TaskRepository) DeleteById(id int) {
	todo := todo.Todo{}
	tr.db.Where("id = ?", id).Find(&todo)
	tr.db.Delete(&todo)
}
