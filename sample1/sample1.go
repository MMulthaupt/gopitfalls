

package niceslice

func GetSliceOfIntPointers() []*int {
	ints := make([]*int, 10)
	for i := 0; i < 10; i++ {
		ints[i] = &i
	}
	return ints
}
