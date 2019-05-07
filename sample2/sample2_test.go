package greetings_test

import (
	"testing"

	greetings "github.com/MMulthaupt/gopitfalls/sample2"
)

func TestGetHelloWorldSuccess(t *testing.T) {
	text, err := greetings.GetHelloWorld(42)
	if err != nil {
		t.Fatalf("Expected success, but got error: %v.", err)
	}
	if text != "Hello World" {
		t.Fatalf("Expected text \"Hello World\", but got \"%s\".", text)
	}
}

func TestGetHelloWorldFailure(t *testing.T) {
	_, err := greetings.GetHelloWorld(0)
	if err == nil {
		t.Fatalf("Expected error, but got nil.")
	}
}
