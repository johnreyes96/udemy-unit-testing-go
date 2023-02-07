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
	amount := calculator.Calculate(50)

	if 50 != amount {
		t.Fail()
	}
}
