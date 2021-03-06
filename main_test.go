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
	"reflect"
	"testing"
)

var (
	repo *git.Repository
	err  error
)

func setup() {
	// log.Printf("testprint")
	repo, err = git.OpenRepository("./test")
	if err != nil {
		log.Println("open repository error: ", err)
		repo, err = git.Clone("https://github.com/Choestelus/vimrc.git", "./test", &git.CloneOptions{})
		if err != nil {
			log.Panicln(err)
		}
	}

	head, err := repo.Head()
	if err != nil {
		log.Fatalln("get repo head error :", err)
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
func TestGit_to_slice(t *testing.T) {
	var somearr [20]byte
	var someslice []byte
	output := git_to_slice(somearr)
	// t.Logf("expected %T got %T\n", someslice, output)
	if reflect.TypeOf(output) != reflect.TypeOf(someslice) {
		t.Errorf("type mismatched:\nexpect []byte got %T", output)
	}
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
	download(tf, "http://127.0.0.1/_vimrc")
	sha1sum := sha1.Sum(tf.Bytes())
	if !bytes.Equal(sha1sum[:], filesum) {
		t.Errorf("expected %x got %x", filesum, sha1sum)
	}
}
func TestHashGit(t *testing.T) {
	odb, err := repo.Odb()
	if err != nil {
		log.Fatalln(err)
	}

	err = odb.ForEach(func(oid *git.Oid) error {
		obj, err := repo.Lookup(oid)
		if err != nil {
			log.Fatalln("Lookup", err)
		}
		switch obj := obj.(type) {
		default:
		case *git.Blob:
			// fmt.Printf("git hash = [%v]", obj.Id())
			// fmt.Printf("file hash %T = [%x]\n", sha1.Sum(([]byte)("eiei")), sha1.Sum(obj.Contents()))
			someobj := obj.Contents()
			var somestr string
			somestr = fmt.Sprintf("blob %v\000", obj.Size())
			// fmt.Printf("[%v]\n", somestr)
			someobj = append([]byte(somestr), someobj...)
			sha1sum := sha1.Sum(someobj)
			if !bytes.Equal(([]byte)(obj.Id()[:]), sha1sum[:]) {
				log.Fatalf("expected %v got %x", obj.Id(), sha1sum)
			}
		}
		return nil
	})
}
func TestGit(t *testing.T) {

}
