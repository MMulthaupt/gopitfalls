`GetSliceOfIntPointers()` is supposed to return a `[]*int` of length 10, where every
element points to an integer with a value equal to the element's index. Running the
test with `go test` however shows that all values are 10. Are these different integers
which are all `10`, or are all pointers actually pointing to the same `10`-integer?
