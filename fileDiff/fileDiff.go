
package fileDiff

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to readFile: %w", err)
	}
	return raw, nil
}

func Diff(a, b string) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(a, b, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
