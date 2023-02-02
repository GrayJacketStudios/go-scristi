package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"go-scristi/pkg/routes"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "80", "Port to serve on")
	r := mux.NewRouter()
	routes.RegisterPostRoutes(r)
	routes.RegisterPagesRoutes(r)
	routes.RegisterStaticRoute(r)

	fmt.Println("Starting server at port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
