package filesystemutil_test

import (
	"testing"

	filesystemutil "github.com/MMulthaupt/gopitfalls/sample3"
)

func TestGetFileContent(t *testing.T) {
	data, err := filesystemutil.GetFileContent()
	if err != nil {
		t.Fatalf("Could not get file content: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("File content is empty.")
	}
}
