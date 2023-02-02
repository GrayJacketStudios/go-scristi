package routes

import (
	"github.com/gorilla/mux"
	"go-scristi/pkg/controllers"
)

var RegisterPagesRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetIndex).Methods("GET")
}
