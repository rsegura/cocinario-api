package usecase

import (
	"github.com/rsegura/cocinario-api.git/pkg/foods"
)

type foodsUseCase struct {
	foodsRepo foods.Repository
}

func NewFoodsUseCase(f foods.Repository) foods.Usecase {
	return &foodsUseCase{
		foodsRepo: f,
	}
}

func (f *foodsUseCase) GetById(id string) (interface{}, error) {
	return f.foodsRepo.GetById(id)
}

func (f *foodsUseCase) Fetch() (interface{}, error) {
	return f.foodsRepo.Fetch()
}
