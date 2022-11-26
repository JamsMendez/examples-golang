package presenter

import "clean-architecture/domain/model"

type UserPresenter interface {
  ResponseUsers(u []*model.User) []*model.User
}
