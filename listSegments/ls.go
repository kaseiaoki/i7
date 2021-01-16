package listSegments

import (
	"fmt"
	"github.com/kaseiaoki/i7/colorPrint"
	// "io/ioutil"
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

func getlist(path string){
	dir, err := os.Open(path)
	if err != nil { 
		fmt.Println(err) 
	}
	defer dir.Close() 

	dirInfo, _ := dir.Stat()
	if ! dirInfo.IsDir() { 
		fmt.Println("no dir")
	}
  
	fns, err := dir.Readdir(-1) 
	if err != nil { 
		fmt.Println(err)
	}

	for _, fns := range fns {
		if fns.IsDir() {
			colorPrint.Out(fns.Name(), "cyan")
		} else {
			colorPrint.Out(fns.Name(), "white")
		}

	}
}
