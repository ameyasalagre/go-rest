package server

import (
	"net/http"
	"routes"
	"github.com/gorilla/mux"
	"utils"
)

/************************.
*	Start the Server
*	Port 8080
************************/
func Start(logger utils.Logger) {
	router := mux.NewRouter().StrictSlash(true)
	routes.IntiateRoutes(router)
	logger.Info("Server Started at localhost:8080")
	http.ListenAndServe(":8080", router)
}
