package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/libgit2/git2go"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	repo *git.Repository
	err  error
)

func setup() {
	// log.SetOutput(ioutil.Discard)
	log.Printf("testprint")
	repo, err = git.OpenRepository("./test")
	if err != nil {
		log.Println(err)
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
	fmt.Fprintf(ioutil.Discard, "%v", head_commit.Id())
}
func teardown() {
}
func TestMain(m *testing.M) {

	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}
func TestGet(t *testing.T) {
	tf := new(bytes.Buffer)
	filesum := make([]byte, 20)
	// filesum := [20]byte{}
	var err error
	filesum, err = hex.DecodeString("26cafd37688f33d4222b349df050c24d6b30d500")
	if err != nil {
		panic(err)
	}
	download(tf, "https://raw.githubusercontent.com/Choestelus/vimrc/master/_vimrc")
	sha1sum := sha1.Sum(tf.Bytes())
	if !bytes.Equal(sha1sum[:], filesum) {
		t.Errorf("expected %x got %x", filesum, sha1sum)
	}
}
func TestHashGit(t *testing.T) {
	t.Logf("commencing hash test")
}
func TestGit(t *testing.T) {

}
