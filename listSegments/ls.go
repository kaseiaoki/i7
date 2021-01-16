package listSegments

import (
	"fmt"
	"github.com/kaseiaoki/i7/colorPrint"
	"log"
	"os"
	"path/filepath"
)

func Path(path string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	target := filepath.FromSlash(dir + "\\" + path)
	list := getlist(target)
	sortFile(list)
}

func getlist(path string) (fns []os.FileInfo) {
	dir, err := os.Open(path)
	if err != nil { 
		fmt.Println(err) 
	}
	defer dir.Close() 

	dirInfo, _ := dir.Stat()
	if ! dirInfo.IsDir() { 
		fmt.Println("no dir")
	}
  
	fns, err = dir.Readdir(-1) 
	if err != nil { 
		fmt.Println(err)
	}

	return fns
}

type FileInfos []os.FileInfo
type fi struct{ FileInfos }

func sortFile(fns []os.FileInfo) {
	for _, fns := range fns {
		if fns.IsDir() {
			colorPrint.Out(fns.Name(), "cyan")
		} else {
			defer colorPrint.Out(fns.Name(), "white")
		}
	}
}