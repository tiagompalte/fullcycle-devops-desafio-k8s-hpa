package main

import "testing"

func TestGreeting(t *testing.T) {
	resultado := greeting("Code.education Rocks!")
	if resultado != "<b>Code.education Rocks!</b>" {
		t.Errorf("greeting('Code.education Rocks!') \n got: %v \n want: %v", resultado, "<b>Code.education Rocks!</b>")
	}
}
