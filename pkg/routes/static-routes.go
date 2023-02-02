package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type NoListFile struct {
	http.File
}

func (f NoListFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

type NoListFileSystem struct {
	base http.FileSystem
}

func (fs NoListFileSystem) Open(name string) (http.File, error) {
	f, err := fs.base.Open(name)
	if err != nil {
		return nil, err
	}
	return NoListFile{f}, nil
}

func intercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var RegisterStaticRoute = func(router *mux.Router) {
	fs := http.FileServer(NoListFileSystem{http.Dir("./static")})
	router.PathPrefix("/public").Handler(http.StripPrefix("/public/", intercept(fs))).Methods("GET")
}
