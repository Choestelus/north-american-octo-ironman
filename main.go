package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/pat"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func download(out io.Writer, url string) {
	tokens := strings.Split(url, "/")
	filename := tokens[len(tokens)-1]
	log.Println("Downloading", url, "to", filename)
	fmt.Println(os.Args)
	fileout := out
	//now downloading via http.Get()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading", url, ": ", err)
		return
	}
	defer resp.Body.Close()

	//TODO: redirect progress to progress bar
	n, err := io.Copy(fileout, resp.Body)
	if err != nil {
		fmt.Println("error while io.Copy() operation : ", err)
		return
	}

	fmt.Println(n, "bytes downloaded")

}

func Hello() int {
	return 42
}

func main() {
	download(os.Stdout, "https://raw.githubusercontent.com/drewolson/testflight/master/README.md")
	// boilerplating import

	// equivalent to router.NewRoute().Path("/").HandlerFunc(HomeHandler)
	// it's all go down to Route type

	// TODO: encapsulate http.Handler in HandlerFunc
	// Here there are 2 parts.
	// 1 is register route path
	// 2 is to register http.Handler to the route

}
