package repository

import (
	"clean-architecture/domain/model"
	"database/sql"
)

type userRepository struct {
  db *sql.DB
}

type UserRepository interface {
  FindAll(users []*model.User) ([]*model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
  return &userRepository{db}
}

func (uR userRepository) FindAll(users []*model.User) ([]*model.User, error) {
  // SQL query using uR.db
  // insert rows in users 
  var err error
  if err != nil {
    return nil, err
  }

  return users, err
}
