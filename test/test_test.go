package test

import (
	"testing"
)

func EchoSolver(s string) string {
	return s
}

func TestSolve(t *testing.T) {
	err := SolveTestfile("echo_tiny.in", EchoSolver)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test succeeded")
	}
}
