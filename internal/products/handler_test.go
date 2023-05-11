package products

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlerGetProductsSuccessful(t *testing.T) {
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

	serviceMock := new(ServiceMock)
	serviceMock.On("GetAllBySeller", "FEX112AC").Return(prodList, nil)

	expectedHttpStatusCode := http.StatusOK
	expectedHttpHeaders := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}
	expectedResponse := `[{
	    "ID": "mock",
	    "SellerID": "FEX112AC",
	    "Description": "generic product",
	    "Price": 123.55
	}]`
	seller_id := "FEX112AC"
	engine := gin.Default()
	handler := NewHandler(serviceMock)
	engine.GET("/api/v1/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/api/v1/products?seller_id="+seller_id, nil)
	resp := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, expectedHttpStatusCode, resp.Code)
	assert.Equal(t, expectedHttpHeaders, resp.Header())
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}

func TestHandlerGetProductsErrSellerIDParamMissing(t *testing.T) {
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

	serviceMock := new(ServiceMock)
	serviceMock.On("GetAllBySeller", "FEX112AC").Return(prodList, nil)

	expectedHttpStatusCode := http.StatusBadRequest
	expectedHttpHeaders := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}
	expectedResponse := `{
		"error": "seller_id query param is required"
	}`
	engine := gin.Default()
	handler := NewHandler(serviceMock)
	engine.GET("/api/v1/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	resp := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, expectedHttpStatusCode, resp.Code)
	assert.Equal(t, expectedHttpHeaders, resp.Header())
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}

func TestHandlerGetProductsErrInternalServer(t *testing.T) {
	// Arrange
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "XXXX",
		Description: "generic product",
		Price:       123.55,
	})
	repoMock := new(RepositoryMock)
	repoMock.On("GetAllBySeller", "XXXX").Return(prodList, nil)

	serviceMock := new(ServiceMock)
	errService := errors.New("internal server error")
	serviceMock.On("GetAllBySeller", "XXXX").Return([]Product(nil), errService)

	expectedHttpStatusCode := http.StatusInternalServerError
	expectedHttpHeaders := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}
	expectedResponse := `{"error":"internal server error"}`
	seller_id := "XXXX"
	engine := gin.Default()
	handler := NewHandler(serviceMock)
	engine.GET("/api/v1/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/api/v1/products?seller_id="+seller_id, nil)
	resp := httptest.NewRecorder()

	// Act
	engine.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, expectedHttpStatusCode, resp.Code)
	assert.Equal(t, expectedHttpHeaders, resp.Header())
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}
