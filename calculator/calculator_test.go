package calculator

import "testing"

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, purchaseAmount: 350, expectedAmount: 290},
		{name: "should not apply", minimumPurchaseAmount: 100, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				// FailNow + log
				t.Fatalf("could not instantiate the calculator %s", err.Error())
			}

			amount := calculator.Calculate(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}
		})
	}
}

func TestDiscountCalculatorShouldFailWithZeroPurchaseAmount(t *testing.T) {
	discountRepositoryMock := DiscountRepositoryMock{}
	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		// FailNow + log
		t.Fatalf("should not create the calculator with zero purchase amount")
	}
}
