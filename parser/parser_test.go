package parser

import "testing"

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
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
	expected := []string{"[Section]", "   C   D   E", "Foo lyric lyric"}
	parser := ParsedContent{}
	err := parser.ParseContent(content)
	if err != nil {
		t.Errorf("Error parsing content: %s", err)
	}
	if !stringSlicesEqual(parser.Lines, expected) {
		t.Errorf("Expected:\n'%#v', got:\n'%#v'",
			expected, parser.Lines)
	}
}
