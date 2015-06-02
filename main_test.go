package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

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
