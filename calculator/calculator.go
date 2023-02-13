package calculator

import "errors"

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount, discountAmount int) (*DiscountCalculator, error) {
	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not bet zero")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}, nil
}

func (c DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {
		multiplier := purchaseAmount / c.minimumPurchaseAmount
		return purchaseAmount - (c.discountAmount * multiplier)
	}

	return purchaseAmount
}
