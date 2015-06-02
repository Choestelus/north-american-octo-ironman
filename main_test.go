package main

import (
	// "github.com/drewolson/testflight"
	"bytes"
	"crypto/sha1"
	// "os"
	"testing"
)

func TestHello(t *testing.T) {
	fn := Hello
	num := fn()
	if num != 42 {
		t.Error("expected 42 got", num)
	}
}
func TestSimpleResponse(t *testing.T) {
	t.Logf("lorem ipsum")
}
func TestGet(t *testing.T) {
	tf := new(bytes.Buffer)
	download(tf, "https://raw.githubusercontent.com/drewolson/testflight/master/README.md")
	t.Logf("sha1sum :%x\n", sha1.Sum(tf.Bytes()))
}
