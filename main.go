package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/parinz1234/todo-go/auth"
	"github.com/parinz1234/todo-go/todo"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tokenz", auth.AccessToken)

	handler := todo.NewTodoHandler(db)

	r.POST("/todos", handler.NewTask)

	r.Run()

	// Can change PORT with additional ENV at runtime ex. PORT=8080 go run main.go
}
