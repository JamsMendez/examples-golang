package controllers

import (
	"clean-architecture/domain/model"
	"clean-architecture/usercase/interactor"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
  userInteractor interactor.UserInteractor
}

type UserController interface {
  GetUsers(c echo.Context) error
}

func NewUserController(uI interactor.UserInteractor) UserController {
  return &userController{uI}
}

func (uC userController) GetUsers(c echo.Context) error {
  var u []*model.User

  u, err := uC.userInteractor.Get(u)
  if err != nil {
    return err
  }

	return c.JSON(http.StatusOK, u)
}
