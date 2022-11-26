package registry

import (
	"clean-architecture/interface/controllers"
	"database/sql"
)

type registry struct {
  db *sql.DB
}

type Registry interface {
  NewAppController() controllers.AppController
}

func NewRegistry(db *sql.DB) Registry {
  return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
  return r.NewUserController()
}
