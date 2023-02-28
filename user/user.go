package user

import (
	"errors"
	"strings"
	"udemy-unit-testing-go/database"

	"udemy-unit-testing-go/entity"
)

type UserService struct {
	userService        database.User
	badWordsRepository database.BadWords
}

func NewUserService(userService database.User, badWordsRepository database.BadWords) *UserService {
	return &UserService{
		userService:        userService,
		badWordsRepository: badWordsRepository,
	}
}

func (c *UserService) Register(user entity.User) error {
	badWords, err := c.badWordsRepository.findAll()
	if err != {
		return err
	}

	for _, badWord := range badWords {
		if strings.Contains(badWord, user.Description) {
			return errors.New("bad word found")
		}
	}

	err = c.userService.Add(user)
	if err != nil {
		return err
	}

	return nil
}
