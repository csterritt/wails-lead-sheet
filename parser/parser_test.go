package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/samber/lo"
)

const content = `
[Section]   
   C   D   E   
Foo lyric lyric
a - B|C / / /| D E
`

func TestImportContent(t *testing.T) {
	expected := lo.Map([]string{"", "[Section]", "   C   D   E", "Foo lyric lyric", "a - B|C / / /| D E", ""}, func(s string, _ int) Line {
		return Line{Text: s, Type: LineTypes.TEXT}
	})
	parser := ParsedContent{}
	err := parser.importContent(content)
	if err != nil {
		t.Errorf("Error parsing content: %s", err)
	}

	if !reflect.DeepEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func TestAllAreChords(t *testing.T) {
	if !allAreChords(makeLetterRuns("A D G")) {
		t.Errorf("Chords found to not be chords")
	}

	if !allAreChords(makeLetterRuns("A#M DbMAJ7b5 GDim")) {
		t.Errorf("Chords with capitalized colors found to not be chords")
	}

	if !allAreChords(makeLetterRuns("A/C Db/Gb GDim/C")) {
		t.Errorf("Chords with inversions found to not be chords")
	}

	if !allAreChords(makeLetterRuns("N.C.")) {
		t.Errorf("N.C. marks found to not be chords")
	}

	for _, suffix := range strings.Split(chordSuffixes, " ") {
		arr := lo.Map([]string{"A", "C#", "Gb"}, func(s string, _ int) string {
			return s + suffix
		})
		if !allAreChords(makeLetterRuns(strings.Join(arr, " "))) {
			t.Errorf("Not all %s chords are chords", suffix)
		}
	}

	if allAreChords(makeLetterRuns("Foo lyric dude")) {
		t.Errorf("Non-chords found to be chords")
	}
}

func TestCategorizeLines(t *testing.T) {
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
		{Text: "", Type: LineTypes.EMPTY},
		{Text: "[Section]", Type: LineTypes.SECTION},
		{Text: "   C   D   E", Type: LineTypes.CHORDS},
		{Text: "Foo lyric lyric", Type: LineTypes.LYRICS},
		{Text: "a - B|C / / /| D E", Type: LineTypes.CHORDS},
		{Text: "", Type: LineTypes.EMPTY},
	}
	if !reflect.DeepEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func TestCategorizeLinesWithNCs(t *testing.T) {
	parser := ParsedContent{}
	contentWithNCs := `
[Section]   
   C   D   E   
Foo lyric lyric
N.C.   N.C.
Spoken line
`

	err := parser.importContent(contentWithNCs)
	if err != nil {
		t.Errorf("Error parsing content with N.C.s: %s", err)
	}

	err = parser.categorizeLines()
	if err != nil {
		t.Errorf("Error categorizing lines: %s", err)
	}

	expected := []Line{
		{Text: "", Type: LineTypes.EMPTY},
		{Text: "[Section]", Type: LineTypes.SECTION},
		{Text: "   C   D   E", Type: LineTypes.CHORDS},
		{Text: "Foo lyric lyric", Type: LineTypes.LYRICS},
		{Text: "N.C.   N.C.", Type: LineTypes.CHORDS},
		{Text: "Spoken line", Type: LineTypes.LYRICS},
		{Text: "", Type: LineTypes.EMPTY},
	}
	if !reflect.DeepEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func TestCategorizeLinesSharpChords(t *testing.T) {
	parser := ParsedContent{}
	contentWithSharpChords := `
[Section]   
   C   D   E   
Foo lyric lyric
C#m7       Asus2/C#        C#m7
Line with sharp chords
`

	err := parser.importContent(contentWithSharpChords)
	if err != nil {
		t.Errorf("Error importing content with sharp chords: %s", err)
	}

	err = parser.categorizeLines()
	if err != nil {
		t.Errorf("Error categorizing lines: %s", err)
	}

	expected := []Line{
		{Text: "", Type: LineTypes.EMPTY},
		{Text: "[Section]", Type: LineTypes.SECTION},
		{Text: "   C   D   E", Type: LineTypes.CHORDS},
		{Text: "Foo lyric lyric", Type: LineTypes.LYRICS},
		{Text: "C#m7       Asus2/C#        C#m7", Type: LineTypes.CHORDS},
		{Text: "Line with sharp chords", Type: LineTypes.LYRICS},
		{Text: "", Type: LineTypes.EMPTY},
	}
	if !reflect.DeepEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}

func TestCompactLines(t *testing.T) {
	compactContent := `

[Section]   


   C   D   E   

Foo lyric lyric

`
	parser := ParsedContent{}
	err := parser.importContent(compactContent)
	if err != nil {
		t.Errorf("Error parsing content: %s", err)
	}

	err = parser.categorizeLines()
	if err != nil {
		t.Errorf("Error categorizing lines: %s", err)
	}

	err = parser.compactLines()
	if err != nil {
		t.Errorf("Error compacting lines: %s", err)
	}

	expected := []Line{
		{Text: "[Section]", Type: LineTypes.SECTION, LineNumber: 0},
		{Text: "", Type: LineTypes.EMPTY, LineNumber: 1},
		{Text: "   C   D   E", Type: LineTypes.CHORDS, LineNumber: 2},
		{Text: "", Type: LineTypes.EMPTY, LineNumber: 3},
		{Text: "Foo lyric lyric", Type: LineTypes.LYRICS, LineNumber: 4},
	}
	if !reflect.DeepEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n")
		for _, line := range expected {
			t.Errorf("  %#v\n", line)
		}
		t.Errorf("Got:\n")
		for _, line := range parser.Lines {
			t.Errorf("  %#v\n", line)
		}
	}
}

func TestMakeLetterRuns(t *testing.T) {
	parts := makeLetterRuns("A B C")
	expected := []LetterRun{
		{Letters: "A", Type: LetterRunTypes.CHORDRUN},
		{Letters: " ", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "B", Type: LetterRunTypes.CHORDRUN},
		{Letters: " ", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "C", Type: LetterRunTypes.CHORDRUN},
	}

	if !reflect.DeepEqual(parts, expected) {
		t.Errorf("Expected:\n'%#v'\ngot:\n'%#v'", expected, parts)
	}

	parts = makeLetterRuns("A#m7b5 - BbDIM/F#|//|CmaJ7")
	expected = []LetterRun{
		{Letters: "A#m7b5", Type: LetterRunTypes.CHORDRUN},
		{Letters: " - ", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "BbDIM/F#", Type: LetterRunTypes.CHORDRUN},
		{Letters: "|//|", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "CmaJ7", Type: LetterRunTypes.CHORDRUN},
	}

	if !reflect.DeepEqual(parts, expected) {
		t.Errorf("Expected:\n'%#v'\ngot:\n'%#v'", expected, parts)
	}

	parts = makeLetterRuns("These abcdefgre lyrics")
	expected = []LetterRun{
		{Letters: "These", Type: LetterRunTypes.WORDRUN},
		{Letters: " ", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "abcdefgre", Type: LetterRunTypes.WORDRUN},
		{Letters: " ", Type: LetterRunTypes.SEPARATORRUN},
		{Letters: "lyrics", Type: LetterRunTypes.WORDRUN},
	}

	if !reflect.DeepEqual(parts, expected) {
		t.Errorf("Expected:\n'%#v'\ngot:\n'%#v'", expected, parts)
	}
}
