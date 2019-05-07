package retrospectively_test

import (
	"testing"

	retrospectively "github.com/MMulthaupt/gopitfalls/sample4"
)

func TestDoSomething(t *testing.T) {
	err := retrospectively.DoSomething()
	if err != nil {
		t.Fatal(err)
	}
}
