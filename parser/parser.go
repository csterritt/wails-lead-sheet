package parser

import (
	"strings"
	"unicode"

	"github.com/samber/lo"
)

const chordSuffixes = "m 7 5 dim dim7 aug sus2 sus4 maj7 m7 7sus4 maj9 maj11 maj13 maj9#11 maj13#11 add9 6add9 maj7b5 maj7#5 m6 m9 m11 m13 madd9 m6add9 mmaj7 mmaj9 m7b5 m7#5 6 9 11 13 7b5 7#5 7b9 7"

var knownChordSuffixes map[string]bool

type Line struct {
	Text string
	Type LineType
}

type ParsedContent struct {
	Lines []Line
}

func init() {
	knownChordSuffixes = make(map[string]bool)
	for _, suffix := range strings.Split(chordSuffixes, " ") {
		knownChordSuffixes[suffix] = true
	}
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

	start := 1
	if s[1] == 'b' || s[1] == '#' {
		start = 2
	}

	if len(s) == 2 {
		return true
	}

	_, found := knownChordSuffixes[s[start:]]

	return found
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

		parts := lo.Filter(strings.Split(p.Lines[index].Text, " "), func(s string, _ int) bool {
			return len(s) > 0
		})

		if allAreChords(parts) {
			p.Lines[index].Type = LineTypes.CHORDS
		} else {
			p.Lines[index].Type = LineTypes.LYRICS
		}
	}
	return nil
}
