package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"coopdloop/ecommerce-api/models"
)

func TestProductsHandler(t *testing.T) {
	// Clear products before each test
	products = []models.Product{}

	t.Run("Get Products - Empty", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProductsHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var returnedProducts []models.Product
		err = json.Unmarshal(rr.Body.Bytes(), &returnedProducts)
		if err != nil {
			t.Errorf("Error unmarshaling response body: %v", err)
		}

		if !reflect.DeepEqual(returnedProducts, []models.Product{}) {
			t.Errorf("handler returned unexpected body: got %v want %v", returnedProducts, []models.Product{})
		}
	})

	t.Run("Create Product", func(t *testing.T) {
		product := models.Product{Name: "Test Product", Price: 9.99}
		jsonProduct, _ := json.Marshal(product)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonProduct))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProductsHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var createdProduct models.Product
		json.Unmarshal(rr.Body.Bytes(), &createdProduct)

		if createdProduct.ID != 1 || createdProduct.Name != "Test Product" || createdProduct.Price != 9.99 {
			t.Errorf("handler returned unexpected product: %v", createdProduct)
		}
	})
}

func TestProductHandler(t *testing.T) {
	// Clear products before each test
	products = []models.Product{
		{ID: 1, Name: "Existing Product", Price: 19.99},
	}

	t.Run("Get Product", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/products/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProductHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var returnedProduct models.Product
		json.Unmarshal(rr.Body.Bytes(), &returnedProduct)

		if returnedProduct.ID != 1 || returnedProduct.Name != "Existing Product" || returnedProduct.Price != 19.99 {
			t.Errorf("handler returned unexpected product: %v", returnedProduct)
		}
	})

	t.Run("Update Product", func(t *testing.T) {
		updatedProduct := models.Product{Name: "Updated Product", Price: 29.99}
		jsonProduct, _ := json.Marshal(updatedProduct)

		req, err := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonProduct))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProductHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var returnedProduct models.Product
		json.Unmarshal(rr.Body.Bytes(), &returnedProduct)

		if returnedProduct.ID != 1 || returnedProduct.Name != "Updated Product" || returnedProduct.Price != 29.99 {
			t.Errorf("handler returned unexpected product: %v", returnedProduct)
		}
	})

	t.Run("Delete Product", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/products/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProductHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		}

		if len(products) != 0 {
			t.Errorf("product was not deleted: %v", products)
		}
	})
}
