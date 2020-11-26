package usecase

import (
	"github.com/rsegura/cocinario-api.git/pkg/recipes"
)

type recipesUseCase struct {
	recipesRepo recipes.Repository
}

func NewRecipesUseCase(r recipes.Repository) recipes.Usecase {
	return &recipesUseCase{
		recipesRepo: r,
	}
}

func (r *recipesUseCase) Fetch() (interface{}, error) {
	return r.recipesRepo.Fetch()
}
