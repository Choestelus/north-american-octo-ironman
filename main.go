package main

import (
	"crypto/sha1"
	"fmt"
	//"github.com/libgit2/git2go"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
	"strings"
)

func download(out io.Writer, url string) {
	tokens := strings.Split(url, "/")
	filename := tokens[len(tokens)-1]
	fmt.Fprintf(ioutil.Discard, url, filename)
	//log.Println("Downloading", url, "to", filename)
	fmt.Fprintln(ioutil.Discard, os.Args)
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
	fmt.Fprintln(ioutil.Discard, n, "bytes downloaded")
}
func git_to_slice(arr [20]byte) []byte {
	return arr[:]
}
func gitchecksum(data []byte) []byte {
	beginstr := fmt.Sprintf("blob %v\x00", len(data))
	gitdata := append([]byte(beginstr), data...)
	ret := sha1.Sum(gitdata)
	return ret[:]
}

func main() {
}
