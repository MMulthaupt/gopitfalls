package service_test

import (
	"testing"

	service "github.com/MMulthaupt/gopitfalls/sample6"
)

func TestCallService(t *testing.T) {
	result, err := service.GetResultFromService(42)
	if err != nil {
		t.Fatalf("Got error from service: %v", err)
	}
	if result != "Glorious result for object with id 42." {
		t.Fatalf(`Got result "%s". Expected: "Glorious result for object with id 42."`, result)
	}
}
