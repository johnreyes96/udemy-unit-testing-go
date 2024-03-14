package user

import (
	"udemy-unit-testing-go/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryStub struct {
	mock.Mock
}

func (r *UserRepositoryStub) Add(user entity.User) error {
	args := r.Called(user)

	return args.Error(0)
}

type BadWordsRepositoryStub struct {
	mock.Mock
}

func (r *BadWordsRepositoryStub) FindAll() ([]string, error) {
	args := r.Called()

	return args.Get(0).([]string), args.Error(1)
}
