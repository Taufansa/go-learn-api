package repository

import (
	"go-learn-api/src/modules/people/model"
)

type PeopleRepository interface {
	Save(*model.People) error
	Update(int, *model.People) error
	Delete(int) error
	FindById(int) (*model.People, error)
	FindAll() (*model.Peoples, error)
}
