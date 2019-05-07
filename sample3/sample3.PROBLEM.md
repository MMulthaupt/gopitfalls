`GetFileContent()` is supposed to read the content of this markdown file and return it.
Running the test with `go test` reports that it works, but there is a bug hidden in the
code. What happens if `ioutil.ReadAll()` returns a non-nil error?
