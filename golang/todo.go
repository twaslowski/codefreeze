package main

import "time"

type Todo struct {
	Id        int32     `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"dueDate"`
}

// todo refactor to type alias for map<int, Todo>
type TodoManager struct {
	Todos []Todo
}

func NewTodoManager() TodoManager {
	return TodoManager{[]Todo{}}
}

type TodoService struct {
	TodoManager TodoManager
}

func (t *TodoManager) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

func (t *TodoManager) Complete(id int32) {
	for index, item := range t.Todos {
		if item.Id == id {
			t.Todos[index].Completed = true
		}
	}
}
