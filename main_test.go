package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	message := Run()
	if message != "Hello, Travis!" {
		t.Error("Wrong message!")
	}
}
