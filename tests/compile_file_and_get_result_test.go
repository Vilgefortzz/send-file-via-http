package tests

import (
	"testing"
)

func TestCompileFileAndGetResult1(t *testing.T) {
	out := compileFileAndGetResult("hello.go")
	if out != "Hello, 世界\n" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "Hello, 世界\n")
	}
}

func TestCompileFileAndGetResult2(t *testing.T) {
	out := compileFileAndGetResult("")
	if out != "compiling error" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "compiling error")
	}
}

func TestCompileFileAndGetResult3(t *testing.T) {
	out := compileFileAndGetResult("abcdefg")
	if out != "compiling error" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "compiling error")
	}
}
