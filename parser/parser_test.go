package parser

import (
	"testing"

	"github.com/samber/lo"
)

const content = `
[Section]   
   C   D   E   
Foo lyric lyric
`

func lineSlicesEqual(a, b []Line) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Type != b[i].Type {
			return false
		}
		if v.Text != b[i].Text {
			return false
		}
	}

	return true
}

func TestImportContent(t *testing.T) {
	expected := lo.Map([]string{"[Section]", "   C   D   E", "Foo lyric lyric"}, func(s string, _ int) Line {
		return Line{Text: s, Type: LineTypes.TEXT}
	})
	parser := ParsedContent{}
	err := parser.importContent(content)
	if err != nil {
		t.Errorf("Error parsing content: %s", err)
	}
	if !lineSlicesEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func xTestCategorizeLines(t *testing.T) {
	parser := ParsedContent{}
	err := parser.importContent(content)
	if err != nil {
		t.Errorf("Error parsing content: %s", err)
	}
	err = parser.categorizeLines()
	if err != nil {
		t.Errorf("Error categorizing lines: %s", err)
	}

	expected := []Line{
		{Text: "[Section]", Type: LineTypes.SECTION},
		{Text: "   C   D   E", Type: LineTypes.CHORDS},
		{Text: "Foo lyric lyric", Type: LineTypes.LYRICS},
	}
	if !lineSlicesEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func TestAllAreChords(t *testing.T) {
	if !allAreChords([]string{"A", "D", "G"}) {
		t.Errorf("Chords found to not be chords")
	}

	if !allAreChords([]string{"Ab", "Bb", "Gb"}) {
		t.Errorf("Flat chords found to not be chords")
	}

	if !allAreChords([]string{"A#", "B#", "G#"}) {
		t.Errorf("Sharp chords found to not be chords")
	}

	if !allAreChords([]string{"A7", "B7", "G7"}) {
		t.Errorf("Dominant chords found to not be chords")
	}

	if !allAreChords([]string{"Ab7", "Bb7", "Gb7"}) {
		t.Errorf("Flat dominant seven chords found to not be chords")
	}

	if !allAreChords([]string{"A#7", "B#7", "G#7"}) {
		t.Errorf("Sharp dominant seven chords found to not be chords")
	}

	if !allAreChords([]string{"A5", "B5", "G5"}) {
		t.Errorf("Five chords found to not be chords")
	}

	if !allAreChords([]string{"Am", "Dm", "Gm"}) {
		t.Errorf("Minor chords found to not be chords")
	}

	if !allAreChords([]string{"Abm", "Bbm", "Gbm"}) {
		t.Errorf("Flat minor chords found to not be chords")
	}

	if !allAreChords([]string{"A#m", "B#m", "G#m"}) {
		t.Errorf("Sharp minor chords found to not be chords")
	}

	if !allAreChords([]string{"Am7", "Bm7", "Gm7"}) {
		t.Errorf("Dominant minor chords found to not be chords")
	}

	if !allAreChords([]string{"Abm7", "Bbm7", "Gbm7"}) {
		t.Errorf("Flat dominant 7 minor chords found to not be chords")
	}

	if !allAreChords([]string{"A#m7", "B#m7", "G#m7"}) {
		t.Errorf("Sharp dominant 7 minor chords found to not be chords")
	}

	if allAreChords([]string{"Foo", "lyric", "lyric"}) {
		t.Errorf("Non-chords found to be chords")
	}
}
