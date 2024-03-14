package payment

import (
	"testing"

	"udemy-unit-testing-go/entity"

	"github.com/stretchr/testify/assert"
)

func TestGivenAUserAndCreditCardWhenIsAuthorizedThenShouldHaveASuccessfullAuthorization(t *testing.T) {
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

func TestGivenAUserAndCreditCardWhenIsNotAuthorizedThenShouldHaveAFailedAuthorization(t *testing.T) {
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

func TestGivenAUserAndCreditCardWhenFailureCountIsMoreThan5RetriesThenShouldHaveAForceFailedAuthorization(t *testing.T) {
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
