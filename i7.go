package main

import (
	"flag"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {

	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Default arguments are two string value that is path for json file.")
		return
	}

	fargs := flag.Args()
	a, err := readfile(filepath.FromSlash(fargs[0]))
	if err != nil {
		log.Fatal(err)
	}
	b, err := readfile(filepath.FromSlash(fargs[1]))
	if err != nil {
		log.Fatal(err)
	}

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(a), string(b), false)

	fmt.Println(dmp.DiffPrettyText(diffs))

}

func readfile(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to readFile: %w", err)
	}
	return raw, nil
}
