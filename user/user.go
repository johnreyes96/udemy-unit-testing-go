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
	badWords, err := c.badWordsRepository.FindAll()
	if err != nil {
		return err
	}

	for _, badWord := range badWords {
		if strings.Contains(user.Description, badWord) {
			return errors.New("bad word found")
		}
	}

	if err = c.userService.Add(user); err != nil {
		return err
	}

	return nil
}
