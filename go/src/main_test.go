package main

import "testing"

func TestGreeting(t *testing.T) {
	greeting := greeting("Code.education Rocks!")
	if greeting != "<b>Code.education Rocks!</b>" {
		t.Error("O texto não é: ", greeting)
	}
}