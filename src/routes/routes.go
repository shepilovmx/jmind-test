package routes

import (
	"github.com/gorilla/mux"
	"jmind-test/src/controllers"
	"jmind-test/src/middlewares"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	blockController := controllers.Block

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.LoggingMiddleware)
	apiRouter.Use(middlewares.ContentTypeMiddleware)

	apiRouter.Handle("/block/{blockNumber}/total", blockController.Perform(blockController.Total)).Methods("GET")

	return r
}
