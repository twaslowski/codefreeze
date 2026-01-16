package main

import "time"

type Todo struct {
	Id        int32
	Title     string
	Completed bool
	DueDate   time.Time
}

type TodoManager struct {
	Todos []Todo
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
