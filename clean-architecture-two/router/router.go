package router

import (
	"clean-architecture/interface/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/users", func(context echo.Context) error {
    return c.GetUsers(context)
  })

  return e
}
