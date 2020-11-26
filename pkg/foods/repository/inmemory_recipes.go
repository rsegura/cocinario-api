package repository

import (
	"github.com/rsegura/cocinario-api.git/pkg/foods"
)

type inmemoryFoodsRepository struct {
	foods map[string]interface{}
}

func NewInmemoryFoodsRepository(foods map[string]interface{}) foods.Repository {
	if foods == nil {
		foods = make(map[string]interface{})
	}

	return &inmemoryFoodsRepository{
		foods: foods,
	}
}

func (r *inmemoryFoodsRepository) Fetch() (interface{}, error) {
	values := make([]interface{}, 0, len(r.foods))

	for _, value := range r.foods {
		values = append(values, value)
	}

	return values, nil
}

func (r *inmemoryFoodsRepository) GetById(Id string) (interface{}, error) {

	return r.foods, nil
}
