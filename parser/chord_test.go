package parser

import "testing"

func TestChordCreation(t *testing.T) {
	c := MakeChord("A")
	if c.String() != "A" {
		t.Errorf("Expected A chord, got %#v", c)
	}
}
