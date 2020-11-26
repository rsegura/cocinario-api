package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsegura/cocinario-api.git/pkg/foods"
)

type FoodsHandler struct {
	FUseCase foods.Usecase
}

func (handler *FoodsHandler) RegisterRouter(r *mux.Router) {
	r.HandleFunc("/foods/{id}", handler.FetchFood).Methods("GET", "OPTIONS")
	r.HandleFunc("/foods", handler.FetchFoods).Methods("GET", "OPTIONS")

}

/*func NewFoodsHandler(r *mux.Router, fu foods.Usecase) Server {
	handler := &FoodsHandler{
		FUseCase: fu,
	}

	a := &server{
		router:  r,
		handler: handler,
	}
	return a
}*/

func (handler *FoodsHandler) FetchFoods(w http.ResponseWriter, r *http.Request) {

	foods, _ := handler.FUseCase.Fetch()
	js, err := json.Marshal(foods)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func (handler *FoodsHandler) FetchFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	foods, fetchError := handler.FUseCase.GetById(id)
	if fetchError != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	fmt.Println(foods)
	js, err := json.Marshal(foods)
	fmt.Println(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

/*func (f *FoodsHandler) GetById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		foods, fetchError := f.FUseCase.GetById(id)
		if fetchError != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		fmt.Println(foods)
		js, err := json.Marshal(foods)
		fmt.Println(js)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return

	})

}

func (f *FoodsHandler) Fetch() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		foods, _ := f.FUseCase.Fetch()
		js, err := json.Marshal(foods)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	})
}*/
