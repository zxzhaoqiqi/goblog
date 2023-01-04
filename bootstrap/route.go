package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/zxzhaoqiqi/goblog/pkg/route"
	"github.com/zxzhaoqiqi/goblog/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()

	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
