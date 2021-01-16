package listSegments

import (
	"fmt"
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
    fileInfos, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, fileInfo := range fileInfos {
        fmt.Println(fileInfo.Name())
    }
}
