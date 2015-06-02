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

func main() {
}
