package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "80", "Port to serve on")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting server at port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
