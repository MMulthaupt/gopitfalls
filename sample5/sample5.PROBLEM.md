`ShortString()` should shorten a string by replacing some interval of characters within it with an ellipsis (`...`),
if it would then be shorter than `leadingCount` + `trailingCount` + the length of the ellipsis. Running the tests
with `go test` reveals that this works for all cases except ones with characters which take up more than 1 byte.
Change the function to support those cases.
