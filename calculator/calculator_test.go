package calculator

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Errorf("expected 130 got %v", amount)
	}
}

func TestDiscountMultipliedByTwoApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(200)

	if amount != 160 {
		t.Errorf("expected 160 got %v", amount)
	}
}

func TestDiscountMultipliedByThreeApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(350)

	if amount != 290 {
		t.Errorf("expected 290 got %v", amount)
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if amount != 50 {
		t.Errorf("expected 50, got %v", amount) // Error = Log + Fail
	}
}
