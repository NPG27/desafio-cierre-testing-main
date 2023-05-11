package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerSuccessful(t *testing.T) {
	// Arrange
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	repoMock := new(RepositoryMock)
	repoMock.On("GetAllBySeller", "FEX112AC").Return(prodList, nil)

	service := NewService(repoMock)

	// Act
	products, err := service.GetAllBySeller("FEX112AC")

	// Assert
	repoMock.AssertCalled(t, "GetAllBySeller", "FEX112AC")
	assert.Nil(t, err)
	assert.Equal(t, prodList, products)
}

func TestGetAllBySellerErrRepoNotFound(t *testing.T) {
	// Arrange
	repoMock := new(RepositoryMock)
	expectedError := errors.New("seller not found based on the given id")
	repoMock.On("GetAllBySeller", "NOEXISTS").Return([]Product(nil), expectedError)

	service := NewService(repoMock)

	// Act
	products, err := service.GetAllBySeller("NOEXISTS")

	// Assert
	repoMock.AssertCalled(t, "GetAllBySeller", "NOEXISTS")
	assert.NotNil(t, err)
	assert.Equal(t, []Product(nil), products)
	assert.ErrorIs(t, err, expectedError)
}

func TestGetAllBySellerErrRepoInternalServer(t *testing.T) {
	// Arrange
	repoMock := new(RepositoryMock)
	expectedError := errors.New("internal server error")
	repoMock.On("GetAllBySeller", "XX").Return([]Product(nil), expectedError)

	service := NewService(repoMock)

	// Act
	products, err := service.GetAllBySeller("XX")

	// Assert
	repoMock.AssertCalled(t, "GetAllBySeller", "XX")
	assert.NotNil(t, err)
	assert.Equal(t, []Product(nil), products)
	assert.ErrorIs(t, err, expectedError)
}
