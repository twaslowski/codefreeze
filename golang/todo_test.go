package main

import (
	"testing"
	"time"

	assert2 "github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	var todos TodoManager
	todo := Todo{
		Id:        1,
		Title:     "My first todo",
		Completed: false,
		DueDate:   time.Now().Add(10 * time.Minute),
	}

	todos.Add(todo)
	assert2.Len(t, todos.Todos, 1)
}

func TestCompleteTodo(t *testing.T) {
	var todos TodoManager
	todo := Todo{
		Id:        1,
		Title:     "My first todo",
		Completed: false,
		DueDate:   time.Now().Add(10 * time.Minute),
	}

	todos.Add(todo)
	todos.Complete(1)

	assert2.True(t, todos.Todos[0].Completed)
}
