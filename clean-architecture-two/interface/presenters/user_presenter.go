package presenters

import (
	"clean-architecture/domain/model"
	"fmt"
)

type userPresenter struct {
}

type UserPresenter interface {
  ResponseUsers(users []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
  return &userPresenter{}
}

func (uP *userPresenter) ResponseUsers(users []*model.User) []*model.User {
  for _, user := range users {
    user.Name = fmt.Sprintf("Mr. %s", user.Name)
  }

  return users
}
