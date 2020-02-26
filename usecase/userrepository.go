package usecase

import "github.com/hukurou-s/user-auth-api-sample/domain"

type UserRepository interface {
	FindByName(string) (domain.User, error)
}
