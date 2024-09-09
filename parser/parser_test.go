package parser

import (
	"strings"
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

	for _, suffix := range strings.Split(chordSuffixes, " ") {
		arr := lo.Map([]string{"A", "C", "G"}, func(s string, _ int) string {
			return s + suffix
		})

		if !allAreChords(arr) {
			t.Errorf("Not all %s chords are chords", suffix)
		}

		arr = lo.Map([]string{"Ab", "Cb", "Gb"}, func(s string, _ int) string {
			return s + suffix
		})

		if !allAreChords(arr) {
			t.Errorf("Not all %s chords are chords", suffix)
		}
		arr = lo.Map([]string{"A#", "C#", "G#"}, func(s string, _ int) string {
			return s + suffix
		})

		if !allAreChords(arr) {
			t.Errorf("Not all %s chords are chords", suffix)
		}
	}

	if allAreChords([]string{"Foo", "lyric", "lyric"}) {
		t.Errorf("Non-chords found to be chords")
	}
}
