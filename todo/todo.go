package todo

type Todo struct {
	ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
