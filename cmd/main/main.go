package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"go-scristi/pkg/routes"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "80", "Port to serve on")

func main() {
	flag.Parse()
	r := mux.NewRouter()
	routes.RegisterPostRoutes(r)
	routes.RegisterPagesRoutes(r)
	routes.RegisterStaticRoute(r)

	fmt.Println("Starting server at port " + *port)
	log.Fatal(http.ListenAndServeTLS(":"+*port, os.Getenv("CERT_CRT"), os.Getenv("CERT_KEY"), r))
}
