`GetMessageFromGoRoutine()` should receive the string `Hello, world!` from a go routine, and then return it.
Running tests with `go test` reports success, but running tests with `go test --race` fails. Why does
this happen? How can you fix it?

Does your solution still work when running `runtime.GOMAXPROCS(1)` beforehand?
