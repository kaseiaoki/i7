package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	flag.Parse()
	// defalt arguments
	fa := os.Args[1]
	fb := os.Args[2]

	if len(fa) <= 0 && len(fb) <= 0 {
		fmt.Println("Default arguments are two string value that is path for json file.")
		return
	}
	// opiton arguments
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		// hex mode flag
		l = fs.Bool("line", false, "line mode")
		h = fs.Bool("hex", false, "hex mode")
	
	)
	fs.Parse(os.Args[3:])

	// file a, file b
	a, err := readfile(filepath.FromSlash(fa))
	if err != nil {
		log.Fatal(err)
	}
	b, err := readfile(filepath.FromSlash(fb))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*h, *l, os.Args[3:])
	diff(a, b, *h, *l)

}

func readfile(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to readFile: %w", err)
	}
	return raw, nil
}

func diff(a []byte, b []byte, h bool, l bool) {
	// casted a and b
	var ca = ""
	var cb = ""

	//hex mode
	if h {
		ca = hex.EncodeToString(a)
		cb = hex.EncodeToString(b)
	} else {
		ca = string(a)
		cb = string(b)
	}

	if l {
		linediff(ca, cb)
	} else {
		filediff(ca, cb)
	}

}

func filediff(a , b string){
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(a, b, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}

func linediff(a , b string)  {
    dmp := diffmatchpatch.New()
    a, b, c := dmp.DiffLinesToChars(a, b)
    diffs := dmp.DiffMain(a, b, false)
    result := dmp.DiffCharsToLines(diffs, c)
    fmt.Println(result)
}