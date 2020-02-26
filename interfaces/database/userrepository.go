package database

import "github.com/hukurou-s/user-auth-api-sample/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindByID(id int) (user domain.User, err error) {
	if result := repo.Where(&user, id); result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (repo *UserRepository) FindByName(name string) (user domain.User, err error) {
	if result := repo.Where("name = ?", name).First(&user); result.Error != nil {
		err = result.Error
		return
	}
	return
}
