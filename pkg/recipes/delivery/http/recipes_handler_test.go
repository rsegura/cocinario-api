package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type serviceMock struct {
	fetch func() (interface{}, error)
}

func (r *serviceMock) Fetch() (interface{}, error) {
	return r.fetch()
}

func TestFetchRecipes(t *testing.T) {

	service := &serviceMock{}

	recipesHandler := RecipesHandler{service}

	t.Run("RecipesFound", func(t *testing.T) {
		service.fetch = func() (interface{}, error) {
			return nil, nil
		}

		req, err := http.NewRequest("GET", "/recipes", nil)

		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		res := httptest.NewRecorder()

		recipesHandler.FetchRecipes(res, req)

		assert.EqualValues(t, http.StatusOK, res.Code)
	})
}
