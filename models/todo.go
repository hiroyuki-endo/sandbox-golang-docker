package models

import "errors"

type Todo struct {
	ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (todo *Todo) Start() error {
	if todo.Status != "Ready" {
		err := errors.New("Todo is not ready")
		return err
	}
	todo.Status = "Doing"
	return nil
}

func (todo *Todo) Done() error {
	if todo.Status != "Doing" {
		err := errors.New("Todo is not doing")
		return err
	}
	todo.Status = "Done"
	return nil
}
