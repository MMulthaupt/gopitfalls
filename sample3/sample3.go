

package filesystemutil

import (
	"io/ioutil"
	"os"
)

func GetFileContent() (data []byte, err error) {
	f, err := os.Open("sample3.CONTEXT.md")
	if err != nil {
		return data, err
	}
	defer func() {
		err = f.Close()
	}()
	return ioutil.ReadAll(f)
}
