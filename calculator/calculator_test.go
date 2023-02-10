package calculator

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, expectedAmount: 160},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, expectedAmount: 290},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)

		if amount != tc.expectedAmount {
			t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
		}
	}
}
