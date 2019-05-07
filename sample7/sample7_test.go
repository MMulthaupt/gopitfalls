package concurrent_test

import (
	"testing"

	concurrent "github.com/MMulthaupt/gopitfalls/sample7"
)

func TestGetMessageFromGoRoutine(t *testing.T) {
	message := concurrent.GetMessageFromGoRoutine()
	if message != "Hello, world!" {
		t.Fatalf(`Got string "%s". Expected "%s".`, message, "Hello, world!")
	}
}
