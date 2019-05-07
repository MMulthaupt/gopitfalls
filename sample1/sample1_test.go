package niceslice_test

import (
	"testing"

	niceslice "github.com/MMulthaupt/gopitfalls/sample1"
)

func TestGetSliceOfIntPointers(t *testing.T) {
	ints := niceslice.GetSliceOfIntPointers()
	if len(ints) != 10 {
		t.Fatalf("Unexpected length %d. Expected 10.", len(ints))
	}
	for i := 0; i < 10; i++ {
		if *ints[i] != i {
			t.Errorf("Unexpected value %d at index %d. Expected %d.", *ints[i], i, i)
		}
	}
}
