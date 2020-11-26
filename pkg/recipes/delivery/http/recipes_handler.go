package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsegura/cocinario-api.git/pkg/recipes"
)

type RecipesHandler struct {
	RUseCase recipes.Usecase
}

func (handler *RecipesHandler) RegisterRouter(r *mux.Router) {
	r.HandleFunc("/recipes", handler.FetchRecipes).Methods("GET", "OPTIONS")
}

func (handler *RecipesHandler) FetchRecipes(w http.ResponseWriter, r *http.Request) {

	recipes, _ := handler.RUseCase.Fetch()
	js, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}
