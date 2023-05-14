package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"udemy-unit-testing-go/entity"
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

func TestShouldHaveASuccessfullAuthorization(t *testing.T) {
	user := entity.User{}
	creditCard := entity.CreditCard{}
	attemptHistory := &AttemptHistory{}
	attemptHistory.On("CountFailures", user).Return(1, nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(true, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, _ := paymentService.IsAuthorized(user, creditCard)

	attemptHistory.AssertNotCalled(t, "IncrementFailure", user)
	assert.True(t, isAuthorized)
}

func TestShouldHaveAFailedAuthorization(t *testing.T) {
	user := entity.User{}
	creditCard := entity.CreditCard{}
	attemptHistory := &AttemptHistory{}
	attemptHistory.On("CountFailures", user).Return(1, nil)
	attemptHistory.On("IncrementFailure", user).Return(nil)
	gateway := &GatewayMock{}
	gateway.On("IsAuthorized", user, creditCard).Return(false, nil)
	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, _ := paymentService.IsAuthorized(user, creditCard)

	attemptHistory.AssertCalled(t, "IncrementFailure", user)
	assert.False(t, isAuthorized)
}

func TestShouldHaveAForceFailedAuthorization(t *testing.T) {
	user := entity.User{}
	creditCard := entity.CreditCard{}
	attemptHistory := &AttemptHistory{}
	attemptHistory.On("CountFailures", user).Return(6, nil)
	gateway := &GatewayMock{}
	paymentService := NewPaymentService(attemptHistory, gateway)

	isAuthorized, _ := paymentService.IsAuthorized(user, creditCard)

	gateway.AssertNotCalled(t, "IsAuthorized", user, creditCard)
	attemptHistory.AssertNotCalled(t, "IncrementFailure", user)
	assert.False(t, isAuthorized)
}
