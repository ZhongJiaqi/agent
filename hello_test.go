package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// Refactored function to make the core logic testable
func printMessage() {
	fmt.Println("hello agent!")
}

func TestPrintMessage(t *testing.T) {
	// Capture the output of the printMessage function
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the printMessage function
	printMessage()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "hello agent!\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
