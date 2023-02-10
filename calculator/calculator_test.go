package calculator

import "testing"

func TestDiscountApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if 130 != amount {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {

	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(60)

	t.Log("Hello!")

	if 50 != amount {
		t.Errorf("expected 50, got %v, failed because the discount was not expected to be applied", amount) // Error = Log + Fail
	}
}
