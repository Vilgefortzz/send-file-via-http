package tests

import (
	"testing"
)

func TestFileExtension1(t *testing.T) {
	ext := getExtensionFromFile("hello.go")
	if ext != ".go" {
		t.Errorf("Extension was incorrect, got: %s, want: %s.", ext, ".go")
	}
}

func TestFileExtension2(t *testing.T) {
	ext := getExtensionFromFile("")
	if ext != "" {
		t.Errorf("Extension was incorrect, got: %s, want: %s.", ext, "")
	}
}

func TestFileExtension3(t *testing.T) {
	ext := getExtensionFromFile("abcdefg")
	if ext != "" {
		t.Errorf("Extension was incorrect, got: %s, want: %s.", ext, "")
	}
}
