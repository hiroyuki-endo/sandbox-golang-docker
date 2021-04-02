package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/microsoft/vscode-remote-try-go/models"
)

type TaskRepository struct {
	db *gorm.DB
}

func CreateTaskRepository(db *gorm.DB) *TaskRepository {
	newTaskRepository := new(TaskRepository)
	newTaskRepository.db = db
	return newTaskRepository
}

func (tr *TaskRepository) All() *[]models.Todo {
	todos := []models.Todo{}
	tr.db.Find(&todos)
	return &todos
}

func (tr *TaskRepository) FindById(id int) *models.Todo {
	todo := models.Todo{}
	tr.db.Where("id = ?", id).Find(&todo)
	return &todo
}

func (tr *TaskRepository) Create(newTodo *models.Todo) int {
	tr.db.Create(newTodo)
	return newTodo.ID
}

func (tr *TaskRepository) Save(newTodo *models.Todo) {
	tr.db.Save(newTodo)
}

func (tr *TaskRepository) DeleteAll() {
	todos := []models.Todo{}
	tr.db.Find(&todos)
	tr.db.Delete(&todos)
}

func (tr *TaskRepository) DeleteById(id int) {
	todo := models.Todo{}
	tr.db.Where("id = ?", id).Find(&todo)
	tr.db.Delete(&todo)
}
