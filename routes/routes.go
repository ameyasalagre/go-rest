package routes

import (
	"controllers"
	"github.com/gorilla/mux"
)

/****************************************.
*	Add New App Routes in below function
*	All Server Routes listed here
*
******************************************/
func IntiateRoutes(routes *mux.Router) {
	routes.HandleFunc("/", controllers.Home)
	routes.HandleFunc("/user", controllers.User)
}
