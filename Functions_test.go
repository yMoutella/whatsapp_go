package main

import "testing"

func TestIsStartingChat(t *testing.T) {

	result := isStartingChat("5521981881050")
	expected := true

	if result != expected {
		t.Errorf("Not expected result - exp:%t  resul:%t", expected, result)
	}

}
