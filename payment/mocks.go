package payment

import (
	"udemy-unit-testing-go/entity"

	"github.com/stretchr/testify/mock"
)

type AttemptHistory struct {
	mock.Mock
}

func (a *AttemptHistory) IncrementFailure(user entity.User) error {
	args := a.Called(user)

	return args.Error(0)
}

func (a *AttemptHistory) CountFailures(user entity.User) (int, error) {
	args := a.Called(user)

	return args.Int(0), args.Error(1)
}

type GatewayMock struct {
	mock.Mock
}

func (gm *GatewayMock) IsAuthorized(user entity.User, creditCard entity.CreditCard) (bool, error) {
	args := gm.Called(user, creditCard)

	return args.Bool(0), args.Error(1)
}

func (gm *GatewayMock) Pay(creditCard entity.CreditCard, amount int) error {
	args := gm.Called(creditCard, amount)

	return args.Error(0)
}
