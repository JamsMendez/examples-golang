package repository

import "clean-architecture/domain/model"

type UserRepository interface {
  FindAll(u []*model.User) ([]*model.User, error)
}
