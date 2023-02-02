package routes

import (
	"github.com/gorilla/mux"
	"go-scristi/pkg/controllers"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
}
