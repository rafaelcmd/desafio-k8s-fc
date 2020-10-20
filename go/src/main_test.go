package main

import "testing"

func TestGreeting(t *testing.T) {
	greeting := greeting("Code.education Rocks!")
	if greeting != "Code.education Rocks!" {
		t.Error("O texto não é: ", greeting)
	}
}