package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(go-sandbox-db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Todo struct {
	ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	db := gormConnect()
	defer db.Close()
	db.AutoMigrate(&Todo{})

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		todos := []Todo{}
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&newTodo)
		c.JSON(http.StatusOK, newTodo.ID)
	})

	router.DELETE("/todos", func(c *gin.Context) {
		todos := []Todo{}
		db.Find(&todos)
		db.Delete(&todos)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		todo := Todo{}
		db.Where("id = ?", id).Find(&todo)
		db.Delete(&todo)
	})

	router.Run(":8080")
}
