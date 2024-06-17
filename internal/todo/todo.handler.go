package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	shared "github.com/ryansheehan/go-todo/internal/shared/helpers"
)

func (m *TodoModule) Home(c echo.Context) error {
	return shared.Render(c, http.StatusOK, Home(m.Todos))
}
