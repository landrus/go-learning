package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	bytes := bytes.NewBufferString("word1 word2 word3 word4")

	expected := 4
	result := count(bytes, false, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	bytes := bytes.NewBufferString("word1 word2 word3 word4\nline 2\nline3")

	expected := 3
	result := count(bytes, true, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	bytes := bytes.NewBufferString("word1\nline 2")

	expected := 12
	result := count(bytes, false, true)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}
