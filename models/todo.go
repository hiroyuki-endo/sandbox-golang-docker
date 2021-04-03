package models

type Todo struct {
	ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (todo *Todo) Done() {
	todo.Status = "Done"
}

func (todo *Todo) Start() {
	todo.Status = "Doing"
}
