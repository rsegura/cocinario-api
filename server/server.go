package server

import (
	"net/http"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/rsegura/cocinario-api.git/pkg/foods"
	_foodHttpDeliver "github.com/rsegura/cocinario-api.git/pkg/foods/delivery/http"
	"github.com/rsegura/cocinario-api.git/pkg/middleware"
	"github.com/rsegura/cocinario-api.git/pkg/recipes"
	_recipeHttpDeliver "github.com/rsegura/cocinario-api.git/pkg/recipes/delivery/http"
)

type Server struct {
	foodService    foods.Usecase
	recipesService recipes.Usecase

	Router *mux.Router
}

func NewServer(foodService foods.Usecase, recipesService recipes.Usecase) *Server {
	server := &Server{
		foodService:    foodService,
		recipesService: recipesService,
	}

	foodHandler := _foodHttpDeliver.FoodsHandler{foodService}
	recipeHandler := _recipeHttpDeliver.RecipesHandler{recipesService}

	router := mux.NewRouter()

	foodHandler.RegisterRouter(router)
	recipeHandler.RegisterRouter(router)
	server.Router = router
	return server
}

func (srv *Server) Run(port string, logger log.Logger) {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         port,
		Handler:      middleware.LoggingMiddleware(logger)(srv.Router),
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.Log(
			"err", err,
		)
	}
}
