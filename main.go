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

// func FileHandler(w http.ResponseWriter, r *http.Request) {
// 	fs := http.FileServer(http.Dir("~/color-spaces.pl"))
// 	return fs
// }

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/upload", UploadHandler)
	router.HandleFunc("/show", ShowHandler)

	// TODO: encapsulate http.Handler in HandlerFunc
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", router)
	http.ListenAndServe(":5001", router)
}
