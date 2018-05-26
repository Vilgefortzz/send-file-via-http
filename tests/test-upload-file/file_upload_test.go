package test_extension_file

import (
	"testing"
)

func TestFileUpload1(t *testing.T) {
	out := upload("hello.go", "localhost")
	expectedOut := "__Output__\n------------\nHello, 世界\n------------\n"
	if out != expectedOut {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, expectedOut)
	}
}

func TestFileUpload2(t *testing.T) {
	out := upload("", "")
	if out != "file not send" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not send")
	}
}

func TestFileUpload3(t *testing.T) {
	out := upload("hello.go", "")
	expectedOut := "__Output__\n------------\nHello, 世界\n------------\n"
	if out != expectedOut {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, expectedOut)
	}
}

func TestFileUpload4(t *testing.T) {
	out := upload("abcdefg", "")
	if out != "file not found" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not found")
	}
}

func TestFileUpload5(t *testing.T) {
	out := upload("", "localhost")
	if out != "file not send" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not send")
	}
}

func TestFileUpload6(t *testing.T) {
	out := upload("", "abcdefg")
	if out != "file not send" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not send")
	}
}

func TestFileUpload7(t *testing.T) {
	out := upload("abcdefg", "localhost")
	if out != "file not found" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not found")
	}
}

func TestFileUpload8(t *testing.T) {
	out := upload("abcdefg", "abcdefg")
	if out != "file not found" {
		t.Errorf("Output was incorrect, got: %s, want: %s.", out, "file not found")
	}
}
