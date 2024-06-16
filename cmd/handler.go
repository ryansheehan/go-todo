package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ryansheehan/go-todo/internal/view/count"
	"github.com/ryansheehan/go-todo/internal/view/page"
)

func (app *application) home(c echo.Context) error {
	return render(c, http.StatusOK, page.Home(&app.state))
}

func (app *application) incrementCount(c echo.Context) error {
	app.state.Count = app.state.Count + 1
	var result = render(c, http.StatusOK, count.Number(app.state.Count))
	return result
}
