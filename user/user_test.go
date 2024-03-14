package user

import (
	"testing"

	"udemy-unit-testing-go/entity"

	"github.com/stretchr/testify/assert"
)

func TestGivenRegisterWhenAddUserIsOkAndNotHasBadWordsThenShouldSuccessfullyRegistrateAnUser(t *testing.T) {
	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software Developer",
	}
	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)
	userService := NewUserService(userRepository, badWordsRepository)

	err := userService.Register(user)

	userRepository.AssertCalled(t, "Add", user)
	assert.Nil(t, err)
}

func TestGivenARegisterWhenBadWordIsFoundThenShouldNotRegistrateTheUser(t *testing.T) {
	user := entity.User{
		Name:        "Vinicius",
		Email:       "vinicius@example.com",
		Description: "Software tomato Developer",
	}
	userRepository := &UserRepositoryStub{}
	userRepository.On("Add", user).Return(nil)
	badWordsRepository := &BadWordsRepositoryStub{}
	badWordsRepository.On("FindAll").Return([]string{"tomato", "potato"}, nil)
	userService := NewUserService(userRepository, badWordsRepository)

	err := userService.Register(user)

	userRepository.AssertNotCalled(t, "Add", user)
	assert.Error(t, err)
}
