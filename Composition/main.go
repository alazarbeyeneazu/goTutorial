package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("The Hash of %s is %s ", p.Path, p.Hash)
}
func (f Pair) FileName() string {
	return filepath.Base(f.Path)
}

type fileNamer interface {
	FileName() string
}

type PairWithLength struct {
	Pair
	Lenght int
}

func main() {
	p := fileNamer(Pair{Path: "hello", Hash: "xslsfhslfdafkasd"})
	fmt.Println(p.FileName())
}
