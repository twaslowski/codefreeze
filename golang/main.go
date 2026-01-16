package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	todoManager := NewTodoManager()
	todoService := TodoService{todoManager}

	r := todoService.setupRouter()
	err := r.Run()
	if err != nil {
		fmt.Printf("Got error when starting webserver: %s", err)
	}
}

func (t TodoService) setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "up",
		})
	})
	r.POST("/todo", t.createTodo)
	r.GET("/todo", t.createTodo)
	r.POST("/todo/%d/complete", t.completeTodo)
	return r
}

func (t TodoService) createTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	t.TodoManager.Add(newTodo)

	context.JSON(201, gin.H{"todo": newTodo})
}

func (t TodoService) findTodo(context *gin.Context) {
	var todo Todo

	context.JSON(201, gin.H{"todo": todo})
}

func (t TodoService) completeTodo(context *gin.Context) {
	var todo Todo
	context.JSON(201, gin.H{"todo": todo})
}
