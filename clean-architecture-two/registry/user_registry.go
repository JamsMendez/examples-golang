package registry

import (
	"clean-architecture/interface/controllers"
	iP "clean-architecture/interface/presenters"
	iR "clean-architecture/interface/repository"
	uP "clean-architecture/usercase/presenter"
	uR "clean-architecture/usercase/repository"
	"clean-architecture/usercase/interactor"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() uR.UserRepository {
	return iR.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() uP.UserPresenter {
	return iP.NewUserPresenter()
}
