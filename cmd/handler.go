package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ryansheehan/go-todo/internal/view/page"
)

func (app *application) home(c echo.Context) error {
	return render(c, http.StatusOK, page.Home("Sam Smith"))

}
