package interactor

import (
	"clean-architecture/domain/model"
	"clean-architecture/usercase/presenter"
	"clean-architecture/usercase/repository"
)

type userInteractor struct {
  UserRepository repository.UserRepository
  UserPresenter presenter.UserPresenter
}

type UserInteractor interface {
  Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
  return &userInteractor{r, p}
}

func (uI *userInteractor) Get(u []*model.User) ([]*model.User, error) {
  u, err := uI.UserRepository.FindAll(u)
  if err != nil {
    return nil, err
  }

  return uI.UserPresenter.ResponseUsers(u), nil
}

