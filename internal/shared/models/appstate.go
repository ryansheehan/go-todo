package shared

import (
	"github.com/ryansheehan/go-todo/internal/todo"
)

type AppState struct {
	Todos []todo.Todo
}

func InitAppState(count uint64, name string) *AppState {
	return &AppState{
		Todos: make([]todo.Todo, 0),
	}
}

type Application struct {
	state AppState
}
