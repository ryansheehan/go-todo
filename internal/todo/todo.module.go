package todo

import (
	"github.com/labstack/echo/v4"
)

type TodoModule struct {
	Todos Todos
}

func NewTodoModule() *TodoModule {
	return &TodoModule{
		Todos: make(Todos, 0),
	}
}

func (m *TodoModule) Mount(e *echo.Echo, baseRoute string) {
	e.GET(baseRoute+"/", m.Home)
}
