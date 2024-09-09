package parser

import (
	"strings"
	"unicode"

	"github.com/samber/lo"
)

type Line struct {
	Text string
	Type LineType
}

type ParsedContent struct {
	Lines []Line
}

func firstNonBlankChar(s string) (rune, bool) {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return c, true
		}
	}

	return ' ', false
}

func isChord(s string) bool {
	if s[0] < 'A' || s[0] > 'G' {
		return false
	}

	if len(s) == 1 {
		return true
	}

	if len(s) == 2 && (s[1] == 'b' || s[1] == '#' || s[1] == '7' || s[1] == '5' || s[1] == 'm') {
		return true
	}

	if len(s) == 3 && (s[1] == 'b' || s[1] == '#') && (s[2] == 'm' || s[2] == '7') {
		return true
	}

	if len(s) == 3 && s[1] == 'm' && s[2] == '7' {
		return true
	}

	if len(s) == 4 && (s[1] == 'b' || s[1] == '#') && s[2] == 'm' && s[3] == '7' {
		return true
	}

	return false
}

func allAreChords(s []string) bool {
	return lo.Reduce(s, func(agg bool, item string, _ int) bool {
		return agg && isChord(item)
	}, true)
}

func (p *ParsedContent) importContent(content string) error {
	p.Lines = lo.FilterMap(strings.Split(content, "\n"), func(s string, _ int) (Line, bool) {
		res := strings.TrimRight(s, " \t\r\n")
		return Line{Text: res, Type: LineTypes.TEXT}, len(res) > 0
	})

	return nil
}

func (p *ParsedContent) categorizeLines() error {
	for index := range p.Lines {
		first, found := firstNonBlankChar(p.Lines[index].Text)
		if found && first == '[' {
			p.Lines[index].Type = LineTypes.SECTION
			continue
		}

		if !found {
			p.Lines[index].Type = LineTypes.EMPTY
			continue
		}

		parts := strings.Split(p.Lines[index].Text, " ")
		if allAreChords(parts) {
			p.Lines[index].Type = LineTypes.CHORDS
		} else {
			p.Lines[index].Type = LineTypes.LYRICS
		}
	}
	return nil
}
