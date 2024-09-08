package parser

import (
	"testing"

	"github.com/samber/lo"
)

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

func TestIgnoreNonHeader(t *testing.T) {
	content := `
[Section]   
   C   D   E   
Foo lyric lyric
`
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
