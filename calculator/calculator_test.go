package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			name:                  "WhenDiscountIs20ThenShouldApply20",
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			name:                  "WhenDiscountIs20AndPurchaseAmountIsDoubleThanMinimumThenShouldApply40",
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			discount:              20,
			expectedAmount:        160,
		},
		{
			name:                  "WhenDiscountIs20AndPurchaseAmountIsTripleThanMinimumThenShouldApply60",
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			discount:              20,
			expectedAmount:        290,
		},
		{
			name:                  "WhenDiscountIs20AndPurchaseAmountIsLessThanMinimumThenShouldNotApply",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              20,
			expectedAmount:        50,
		},
		{
			name:                  "WhenDiscountIsZeroThenShouldNotApply",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
			calculator, _ := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)

			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, tc.expectedAmount, amount)
		})
	}
}

func TestGivenDiscountCalculatorWhenMinimumPurchaseAmountIsZeroThenShouldFailWithZeroPurchaseAmount(t *testing.T) {
	discountRepositoryMock := DiscountRepositoryMock{}

	_, err := NewDiscountCalculator(0, discountRepositoryMock)

	if err == nil {
		t.Fatalf("should not create the calculator with zero purchase amount")
	}
}
