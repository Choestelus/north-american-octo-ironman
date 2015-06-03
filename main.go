package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"io"
	"io/ioutil"
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
	repo, err := git.OpenRepository("./test")
	if err != nil {
		//log.Println(err)
		repo, err = git.Clone("https://github.com/Choestelus/vimrc.git", "./test", &git.CloneOptions{})
		if err != nil {
			log.Panicln(err)
		}
	}
	head, err := repo.Head()
	if err != nil {
		log.Fatalln("error :", err)
	}
	head_commit, err := repo.LookupCommit(head.Target())
	//log.Printf("latest commit id = %v\n", head_commit.Id())
	//log.Printf("Summary: %v\n", head_commit.Summary())
	//fmt.Println("==========================================================")
	odb, err := repo.Odb()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(ioutil.Discard, "%v %v", odb, head_commit.Id())
}
