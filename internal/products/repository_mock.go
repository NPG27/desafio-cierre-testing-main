package products

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}
