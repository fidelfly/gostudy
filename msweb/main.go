package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"flag"
	"time"
	"log"
)

type HomeRouter struct {

}

func (r *HomeRouter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "index.html")
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	//router.Handle("/", &HomeRouter{})

	srv := &http.Server{
		Handler:router,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
