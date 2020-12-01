package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type serviceMock struct {
	getbyId func(id string) (interface{}, error)
	fetch   func() (interface{}, error)
}

func (c *serviceMock) GetById(id string) (interface{}, error) {
	return c.getbyId(id)
}

func (c *serviceMock) Fetch() (interface{}, error) {
	return c.fetch()
}

func TestFetchFood(t *testing.T) {
	service := &serviceMock{}

	foodHandler := FoodsHandler{service}

	t.Run("FoodNotfound", func(t *testing.T) {
		service.getbyId = func(id string) (interface{}, error) {
			return nil, errors.New("food 123 not found")
		}

		req, err := http.NewRequest("GET", "/foods/123", nil)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		res := httptest.NewRecorder()

		foodHandler.FetchFood(res, req)

		assert.EqualValues(t, http.StatusNoContent, res.Code)

	})

	t.Run("Foodfound", func(t *testing.T) {
		service.getbyId = func(id string) (interface{}, error) {
			return nil, nil
		}

		req, err := http.NewRequest("GET", "/foods/123", nil)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		res := httptest.NewRecorder()

		foodHandler.FetchFood(res, req)

		assert.EqualValues(t, http.StatusOK, res.Code)

	})
}

func TestFetchFoods(t *testing.T) {
	service := &serviceMock{}

	foodHandler := FoodsHandler{service}

	t.Run("FoodNotfound", func(t *testing.T) {
		service.getbyId = func(id string) (interface{}, error) {
			return nil, errors.New("foods not found")
		}

		req, err := http.NewRequest("GET", "/foods/", nil)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		res := httptest.NewRecorder()

		foodHandler.FetchFood(res, req)

		assert.EqualValues(t, http.StatusNoContent, res.Code)

	})

	t.Run("Foodfound", func(t *testing.T) {
		service.getbyId = func(id string) (interface{}, error) {
			return nil, nil
		}

		req, err := http.NewRequest("GET", "/foods", nil)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		res := httptest.NewRecorder()

		foodHandler.FetchFood(res, req)

		assert.EqualValues(t, http.StatusOK, res.Code)

	})
}
