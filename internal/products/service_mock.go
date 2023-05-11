package products

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}
