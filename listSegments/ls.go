package listSegments

import (
	"fmt"
	"github.com/kaseiaoki/i7/colorPrint"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var layout = "Mon 01/02/06 15:04:05"

func timeFmt(fns os.FileInfo) string {
	return fns.ModTime().Format(layout)
}

func sizeFMT(fns os.FileInfo) string {
	int64 := strconv.FormatInt(int64(fns.Size()), 10)
	spaces := " ";
	for i := 0; i < (19 - len(int64)); i++ {
		spaces += " ";
	}
	return int64 + spaces
}

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
			colorPrint.Out(fns.ModTime().Format("Mon 01/02/06 15:04:05")  + " "  + fns.Name(), "cyan")
		} else {
			defer colorPrint.Out(timeFmt(fns) + " " + sizeFMT(fns) + " " + fns.Name(), "white")
		}
	}
}

