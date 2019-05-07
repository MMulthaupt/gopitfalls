//+build ignore

package filesystemutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetFileContent() (data []byte, err error) {
	f, err := os.Open("sample2.CONTEXT.md")
	if err != nil {
		return data, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%v: %v", closeErr, err)
			} else {
				err = closeErr
			}
		}
	}()
	return ioutil.ReadAll(f)
}
