package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
	log.Printf("Home route handled\n")
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "upload")
	log.Printf("Upload route handled\n")
}
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "show")
	log.Printf("Show route handled\n")
}
func TestHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
}
func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware called")
		next.ServeHTTP(w, r)
		log.Println("ServeHTTP called")
	})
}

func main() {
	router := mux.NewRouter()
	// equivalent to router.NewRoute().Path("/").HandlerFunc(HomeHandler)
	// it's all go down to Route type
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/upload", UploadHandler)
	router.HandleFunc("/show", ShowHandler)

	something := TestHandler()
	// TODO: encapsulate http.Handler in HandlerFunc
	// Here there are 2 parts.
	// 1 is register route path
	// 2 is to register http.Handler to the route
	router.PathPrefix("/static/").Handler((something))

	http.Handle("/", router)
	http.ListenAndServe(":5001", router)
}
