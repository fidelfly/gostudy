package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/gorilla/mux"
)

type HomeRouter struct {
}

func (r *HomeRouter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("HomeRouter Url %s", req.URL.Path)
	http.ServeFile(rw, req, "index.html")
}

func main() {
	fmt.Println("msweb is running!!!")
	var dir string
	var port string
	flag.StringVar(&dir, "dir", "static", "the directory to serve files from. Defaults to the current dir")
	flag.StringVar(&port, "port", "8090", "The port to listen")
	flag.Parse()
	fmt.Println("Listen on %s", port)
	router := mux.NewRouter()
	router.Handle("/", &HomeRouter{})
	//router.Handle("/home", &HomeRouter{})
	router.PathPrefix("/home").Handler(&HomeRouter{})
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	fmt.Println("msweb is ending!!!")
}
