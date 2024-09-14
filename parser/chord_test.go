package parser

import (
	"strings"
	"testing"
)

func TestChordCreation(t *testing.T) {
	c := MakeChord("")
	if c.String() != "" {
		t.Errorf("Expected Empty chord, got %#v", c)
	}

	c = MakeChord("A")
	if c.String() != "A" {
		t.Errorf("Expected A chord, got %#v", c)
	}

	c = MakeChord("A#")
	if c.String() != "A#" {
		t.Errorf("Expected A# chord, got %#v", c)
	}

	c = MakeChord("Ab")
	if c.String() != "Ab" {
		t.Errorf("Expected Ab chord, got %#v", c)
	}

	c = MakeChord("Am")
	if c.String() != "Am" {
		t.Errorf("Expected Am chord, got %#v", c)
	}

	c = MakeChord("A#m7b5")
	if c.String() != "A#m7b5" {
		t.Errorf("Expected A#m7b5 chord, got %#v", c)
	}

	c = MakeChord("A#m7b5/C#")
	if c.String() != "A#m7b5/C#" {
		t.Errorf("Expected A#m7b5/C# chord, got %#v", c)
	}

	c = MakeChord("Abm6add9")
	if c.String() != "Abm6add9" {
		t.Errorf("Expected Abm6add9 chord, got %#v", c)
	}

	c = MakeChord("N.C.")
	if c.String() != "" {
		t.Errorf("Expected N.C. 'chord' to be empty, got %#v", c)
	}

	c = MakeChord("bogus")
	if c.String() != "" {
		t.Errorf("Expected bogus 'chord' to be empty, got %#v", c)
	}

	c = MakeChord("zowi")
	if c.String() != "" {
		t.Errorf("Expected zowi 'chord' to be empty, got %#v", c)
	}

	for _, suffix := range strings.Split(chordSuffixes, " ") {
		for _, note := range []string{"A", "C#", "Gb"} {
			pattern := note + suffix
			c = MakeChord(pattern)
			if c.String() != pattern {
				t.Errorf("Expected example chord build for %s to succeed, got %#v", pattern, c)
			}

			pattern = note + suffix + "/F#"
			c = MakeChord(pattern)
			if c.String() != pattern {
				t.Errorf("Expected example chord build for %s to succeed, got %#v", pattern, c)
			}
		}
	}
}

func testStepUpChord(t *testing.T, note string, expected string) {
	c := MakeChord(note)
	c.StepUp()
	if c.String() != expected {
		t.Errorf("Expected %s chord from %s, got %s", expected, note, c.String())
	}
}

func testStepDownChord(t *testing.T, note string, expected string) {
	c := MakeChord(note)
	c.StepDown()
	if c.String() != expected {
		t.Errorf("Expected %s chord from %s, got %s", expected, note, c.String())
	}
}

func TestChordStepUp(t *testing.T) {
	testStepUpChord(t, "Ab", "A")
	testStepUpChord(t, "A", "A#")
	testStepUpChord(t, "A#", "B")
	testStepUpChord(t, "Bb", "B")
	testStepUpChord(t, "B", "C")
	testStepUpChord(t, "C", "C#")
	testStepUpChord(t, "C#", "D")
	testStepUpChord(t, "Db", "D")
	testStepUpChord(t, "D", "D#")
	testStepUpChord(t, "D#", "E")
	testStepUpChord(t, "E", "F")
	testStepUpChord(t, "F", "F#")
	testStepUpChord(t, "F#", "G")
	testStepUpChord(t, "Gb", "G")
	testStepUpChord(t, "G", "G#")
	testStepUpChord(t, "G#", "A")

	testStepUpChord(t, "Gbm7b5", "Gm7b5")
	testStepUpChord(t, "Gm", "G#m")
	testStepUpChord(t, "G#sus4", "Asus4")

	testStepUpChord(t, "Bbm7b5/Db", "Bm7b5/D")
	testStepUpChord(t, "Bb/D#", "B/E")
	testStepUpChord(t, "B/D#", "C/E")
}

func TestChordStepDown(t *testing.T) {
	testStepDownChord(t, "Ab", "G")
	testStepDownChord(t, "A", "Ab")
	testStepDownChord(t, "A#", "A")
	testStepDownChord(t, "Bb", "A")
	testStepDownChord(t, "B", "Bb")
	testStepDownChord(t, "C", "B")
	testStepDownChord(t, "C#", "C")
	testStepDownChord(t, "Db", "C")
	testStepDownChord(t, "D", "Db")
	testStepDownChord(t, "D#", "D")
	testStepDownChord(t, "Eb", "D")
	testStepDownChord(t, "E", "Eb")
	testStepDownChord(t, "F", "E")
	testStepDownChord(t, "F#", "F")
	testStepDownChord(t, "Gb", "F")
	testStepDownChord(t, "G", "Gb")
	testStepDownChord(t, "G#", "G")

	testStepDownChord(t, "Gbm7b5", "Fm7b5")
	testStepDownChord(t, "Gm", "Gbm")
	testStepDownChord(t, "G#m", "Gm")

	testStepDownChord(t, "Bm7b5/D", "Bbm7b5/Db")
	testStepDownChord(t, "B/E", "Bb/Eb")
	testStepDownChord(t, "C/E", "B/Eb")
}
