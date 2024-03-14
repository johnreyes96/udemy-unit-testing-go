package calculator

import "github.com/stretchr/testify/mock"

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()
	return args.Int(0)
}
