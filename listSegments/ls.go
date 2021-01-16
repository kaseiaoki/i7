package listSegments

import (
	"fmt"
	"github.com/kaseiaoki/i7/colorPrint"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Current() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	getlist(dir)
}

func Path(path string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	target := filepath.FromSlash(dir + "\\" + path)
	getlist(target)
}

func getlist(dir string) {
	fmt.Println(dir)
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fi {
		if fi.IsDir() {
			colorPrint.Out(fi.Name(), "cyan")
		} else {
			colorPrint.Out(fi.Name(), "white")
		}

	}
}
