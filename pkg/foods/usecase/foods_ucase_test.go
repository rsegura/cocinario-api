package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dbClientMock struct {
	getById func(id string) (interface{}, error)
	fetch   func() (interface{}, error)
}

func (c *dbClientMock) GetById(id string) (interface{}, error) {
	return c.getById(id)
}

func (c *dbClientMock) Fetch() (interface{}, error) {
	return c.fetch()
}

func TestGetFood(t *testing.T) {
	dbMock := dbClientMock{}

	foodsService := NewFoodsUseCase(&dbMock)

	t.Run("ErrorFromDatabase", func(t *testing.T) {
		dbMock.getById = func(id string) (interface{}, error) {
			return nil, errors.New("error connecting to our Database")
		}

		food, err := foodsService.GetById("123")
		assert.Nil(t, food)
		assert.NotNil(t, err)
		assert.EqualValues(t, "error connecting to our Database", err.Error())
	})

	t.Run("FoodNotFound", func(t *testing.T) {
		dbMock.getById = func(id string) (interface{}, error) {
			return nil, errors.New("food 123 not found")
		}

		food, err := foodsService.GetById("123")
		assert.Nil(t, food)
		assert.NotNil(t, err)
		assert.EqualValues(t, "food 123 not found", err.Error())
	})
}
