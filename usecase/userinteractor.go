package usecase

import "github.com/hukurou-s/user-auth-api-sample/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserByName(name string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByName(name)
	return
}
