package parser

import (
	"strings"
	"testing"
)

func compareChordString(t *testing.T, source string, asString string, asOriginal string) {
	c := MakeChord(source)
	if c.String() != asString {
		t.Errorf("Expected %#v String(), got String() %#v", asString, c.String())
	}

	if c.OriginalString != asOriginal {
		t.Errorf("Expected %#v original, got OriginalString %#v", asOriginal, c.OriginalString)
	}
}

func TestChordCreation(t *testing.T) {
	compareChordString(t, "", "", "")
	compareChordString(t, "a", "A", "A")
	compareChordString(t, "a/c#", "A/C#", "A/C#")
	compareChordString(t, "A", "A", "A")
	compareChordString(t, "A#", "A#", "A#")
	compareChordString(t, "Ab", "Ab", "Ab")
	compareChordString(t, "Am", "Am", "Am")
	compareChordString(t, "A#m7b5", "A#m7b5", "A#m7b5")
	compareChordString(t, "A#m7b5/C#", "A#m7b5/C#", "A#m7b5/C#")
	compareChordString(t, "Abm6add9", "Abm6add9", "Abm6add9")
	compareChordString(t, "N.C.", "", "")
	compareChordString(t, "bogus", "", "")
	compareChordString(t, "zowi", "", "")

	for _, suffix := range strings.Split(chordSuffixes, " ") {
		for _, note := range []string{"A", "C#", "Gb"} {
			pattern := note + suffix
			c := MakeChord(pattern)
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

func TestChordStepDownThenUp(t *testing.T) {
	c := MakeChord("A")
	c.StepDown()
	if c.String() != "Ab" {
		t.Errorf("Expected %s chord, got %s", "Ab", c.String())
	}
	c.StepDown()
	if c.String() != "G" {
		t.Errorf("Expected %s chord, got %s", "G", c.String())
	}
	c.StepDown()
	if c.String() != "Gb" {
		t.Errorf("Expected %s chord, got %s", "Gb", c.String())
	}

	c.StepUp()
	if c.String() != "G" {
		t.Errorf("Expected %s chord, got %s", "G", c.String())
	}
	c.StepUp()
	if c.String() != "G#" {
		t.Errorf("Expected %s chord, got %s", "G#", c.String())
	}
	c.StepUp()
	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}
}

func TestChordStepUpThenDown(t *testing.T) {
	c := MakeChord("A")
	c.StepUp()
	if c.String() != "A#" {
		t.Errorf("Expected %s chord, got %s", "A#", c.String())
	}
	c.StepUp()
	if c.String() != "B" {
		t.Errorf("Expected %s chord, got %s", "B", c.String())
	}
	c.StepUp()
	if c.String() != "C" {
		t.Errorf("Expected %s chord, got %s", "C", c.String())
	}

	c.StepDown()
	if c.String() != "B" {
		t.Errorf("Expected %s chord, got %s", "B", c.String())
	}
	c.StepDown()
	if c.String() != "Bb" {
		t.Errorf("Expected %s chord, got %s", "Bb", c.String())
	}
	c.StepDown()
	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}
}

func TestChordStepUpAnOctave(t *testing.T) {
	c := MakeChord("A")

	for range 12 {
		c.StepUp()
	}
	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}
}

func TestChordStepDownAnOctave(t *testing.T) {
	c := MakeChord("A")

	for range 12 {
		c.StepDown()
	}
	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}
}

func TestChordReset(t *testing.T) {
	c := MakeChord("A")
	c.StepDown()
	c.Reset()

	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}

	c.StepUp()
	c.Reset()

	if c.String() != "A" {
		t.Errorf("Expected %s chord, got %s", "A", c.String())
	}
}
