package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juju/mgosession"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"

	_log "github.com/go-kit/kit/log"
	//_foodHttpDeliver "github.com/rsegura/cocinario-api.git/pkg/foods/delivery/http"
	_foodRepo "github.com/rsegura/cocinario-api.git/pkg/foods/repository"
	_foodUseCase "github.com/rsegura/cocinario-api.git/pkg/foods/usecase"
	_recipeRepo "github.com/rsegura/cocinario-api.git/pkg/recipes/repository"
	_recipeUseCase "github.com/rsegura/cocinario-api.git/pkg/recipes/usecase"
	"github.com/rsegura/cocinario-api.git/server"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbName := viper.GetString(`database.name`)
	serverAddress := viper.GetString("server.address")
	session, err := mgo.Dial(dbHost + ":" + dbPort)
	if err != nil {
		fmt.Println("Error al conectar a mongo")
		log.Fatal(err.Error())
	}

	defer session.Close()
	mPool := mgosession.NewPool(nil, session, 10)
	defer mPool.Close()

	logger := _log.NewJSONLogger(os.Stdout)
	logger = _log.With(logger, "ts", _log.DefaultTimestampUTC, "loc", _log.DefaultCaller)

	foodRepo := _foodRepo.NewMongoFoodsRepository(mPool, dbName)
	foodUseCase := _foodUseCase.NewFoodsUseCase(foodRepo)
	recipeRepo := _recipeRepo.NewMongoRecipesRepository(mPool, dbName)
	recipeUseCase := _recipeUseCase.NewRecipesUseCase(recipeRepo)
	srv := server.NewServer(foodUseCase, recipeUseCase)
	srv.Run(serverAddress, logger)

	/*_foodHttpDeliver.NewFoodsHandler(router, foodUseCase)
	logger := _log.NewJSONLogger(os.Stdout)
	logger = _log.With(logger, "ts", _log.DefaultTimestampUTC, "loc", _log.DefaultCaller)
	//loggingHandler := middleware.LoggingMiddleware(logger)

	http.Handle("/", router)
	logger2 := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         viper.GetString("server.address"),
		Handler:      middleware.LoggingMiddleware(logger)(router),
		ErrorLog:     logger2,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}*/

}
