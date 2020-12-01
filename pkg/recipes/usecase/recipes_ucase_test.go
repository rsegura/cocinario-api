package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dbClientMock struct {
	fetch func() (interface{}, error)
}

func (r *dbClientMock) Fetch() (interface{}, error) {
	return r.fetch()
}

func TestGetRecipes(t *testing.T) {
	dbMock := dbClientMock{}

	recipesService := NewRecipesUseCase(&dbMock)

	t.Run("ErrorFromDatabase", func(t *testing.T) {
		dbMock.fetch = func() (interface{}, error) {
			return nil, errors.New("error connecting to our Database")
		}

		recipes, err := recipesService.Fetch()
		assert.Nil(t, recipes)
		assert.NotNil(t, err)
		assert.EqualValues(t, "error connecting to our Database", err.Error())
	})

	t.Run("RecipeNotFound", func(t *testing.T) {
		dbMock.fetch = func() (interface{}, error) {
			return nil, errors.New("Recipes Not Found")
		}

		recipes, err := recipesService.Fetch()
		assert.Nil(t, recipes)
		assert.NotNil(t, err)
		assert.EqualValues(t, "Recipes Not Found", err.Error())
	})
}
