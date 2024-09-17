package parser

import (
	"strings"
	"unicode"

	"github.com/samber/lo"
)

const chordSuffixes = "m 7 5 dim dim7 aug sus sus2 sus4 maj7 m7 7sus4 maj9 maj11 maj13 maj9#11 maj13#11 add9 6add9 maj7b5 maj7#5 m6 m9 m11 m13 madd9 m6add9 mmaj7 mmaj9 m7b5 m7#5 6 9 11 13 7b5 7#5 7b9 7"

type LetterRun struct {
	Type              LetterRunType
	Letters           string
	OriginalLetters   string
	Chord             Chord
	TransposedLetters string
}

type Line struct {
	LineNumber int
	Text       string
	Parts      []LetterRun
	Type       LineType
}

type ParsedContent struct {
	Lines []Line
}

var knownChordSuffixes map[string]bool

var chordLetters map[rune]bool
var separators map[rune]bool

func init() {
	knownChordSuffixes = make(map[string]bool)
	for _, suffix := range strings.Split(chordSuffixes, " ") {
		knownChordSuffixes[suffix] = true
	}

	chordLetters = make(map[rune]bool)
	for _, ch := range chordSuffixes {
		if ch != ' ' {
			chordLetters[ch] = true
		}
	}

	for _, ch := range "abcdefg#n." {
		chordLetters[ch] = true
	}

	separators = make(map[rune]bool)
	for _, ch := range " \t\r\n!\"$%&'()*+,-:;<=>?@[\\]^_`{|}~" {
		if _, found := chordLetters[ch]; !found {
			separators[ch] = true
		}
	}
}

func (line Line) String() string {
	if line.Type == LineTypes.CHORDS {
		res := ""
		for _, part := range line.Parts {
			if part.TransposedLetters != "" {
				res += part.TransposedLetters
			} else {
				res += part.Letters
			}
		}

		return res
	}

	return line.Text
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
	s = strings.ToLower(s)
	found := false
	if s == "n.c." {
		return true
	}

	if strings.Index(s, "/") != -1 {
		parts := strings.Split(s, "/")
		if len(parts) != 2 {
			return false
		}

		return isChord(parts[0]) && isChord(parts[1])
	} else {
		if s[0] < 'a' || s[0] > 'g' {
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

		_, found = knownChordSuffixes[s[start:]]
	}

	return found
}

func allAreChords(s []LetterRun) bool {
	foundOneChord := false
	allAreChordOrSeparators := lo.Reduce(s, func(agg bool, item LetterRun, _ int) bool {
		isAChord := isChord(item.Letters)
		if isAChord {
			foundOneChord = true
		}

		return agg && isAChord || agg && item.Type == LetterRunTypes.SEPARATORRUN
	}, true)

	return foundOneChord && allAreChordOrSeparators
}

func makeOneLetterRun(text string, typ LetterRunType) LetterRun {
	if typ == LetterRunTypes.CHORDRUN {
		return LetterRun{Letters: text, Type: typ, Chord: MakeChord(text)}
	} else {
		return LetterRun{Letters: text, Type: typ}
	}
}

func makeLetterRuns(s string) []LetterRun {
	res := make([]LetterRun, 0)

	currentText := ""
	currentType := LetterRunTypes.UNKNOWNRUN
	for index, ch := range strings.ToLower(s) {
		if _, found := separators[ch]; found {
			if currentType == LetterRunTypes.SEPARATORRUN {
				currentText += string(s[index])
			} else {
				if len(currentText) > 0 {
					res = append(res, makeOneLetterRun(currentText, currentType))
				}
				currentText = string(s[index])
				currentType = LetterRunTypes.SEPARATORRUN
			}
		} else {
			if _, found := chordLetters[ch]; found {
				if currentType == LetterRunTypes.CHORDRUN || currentType == LetterRunTypes.WORDRUN {
					currentText += string(s[index])
				} else {
					if len(currentText) > 0 {
						run := makeOneLetterRun(currentText, currentType)
						if run.Type == LetterRunTypes.SEPARATORRUN {
							run.OriginalLetters = currentText
						}
						res = append(res, run)
					}
					currentText = string(s[index])
					currentType = LetterRunTypes.CHORDRUN
				}
			} else {
				if ch == '/' {
					if len(currentText) == 0 {
						currentType = LetterRunTypes.SEPARATORRUN
					}
					currentText += string(s[index])
				} else {
					if currentType == LetterRunTypes.CHORDRUN || currentType == LetterRunTypes.WORDRUN {
						currentText += string(s[index])
					} else {
						if len(currentText) > 0 {
							run := makeOneLetterRun(currentText, currentType)
							if run.Type == LetterRunTypes.SEPARATORRUN {
								run.OriginalLetters = currentText
							}
							res = append(res, run)
						}
						currentText = string(s[index])
					}
					currentType = LetterRunTypes.WORDRUN
				}
			}
		}
	}

	if len(currentText) > 0 {
		run := makeOneLetterRun(currentText, currentType)
		if run.Type == LetterRunTypes.SEPARATORRUN {
			run.OriginalLetters = currentText
		}
		res = append(res, run)
	}

	return res
}

func (p *ParsedContent) importContent(content string) error {
	p.Lines = lo.Map(strings.Split(content, "\n"), func(s string, _ int) Line {
		res := strings.TrimRight(s, " \t\r\n")
		return Line{Text: res, Type: LineTypes.TEXT}
	})

	return nil
}

func (p *ParsedContent) categorizeLines() error {
	for index := range p.Lines {
		first, found := firstNonBlankChar(p.Lines[index].Text)
		if found && first == '[' {
			p.Lines[index].Type = LineTypes.SECTION
			p.Lines[index].Parts = makeLetterRuns("")
			continue
		}

		if !found {
			p.Lines[index].Type = LineTypes.EMPTY
			p.Lines[index].Parts = makeLetterRuns("")
			continue
		}

		parts := makeLetterRuns(p.Lines[index].Text)
		if allAreChords(parts) {
			p.Lines[index].Type = LineTypes.CHORDS
			p.Lines[index].Parts = parts
		} else {
			p.Lines[index].Type = LineTypes.LYRICS
			p.Lines[index].Parts = makeLetterRuns("")
		}
	}
	return nil
}

func (p *ParsedContent) compactLines() error {
	lastWasEmpty := true
	p.Lines = lo.Filter(p.Lines, func(item Line, index int) bool {
		if item.Type == LineTypes.EMPTY && lastWasEmpty {
			return false
		}

		lastWasEmpty = item.Type == LineTypes.EMPTY
		return true
	})

	if p.Lines[len(p.Lines)-1].Type == LineTypes.EMPTY {
		p.Lines = p.Lines[:len(p.Lines)-1]
	}

	next := 0
	p.Lines = lo.Map(p.Lines, func(item Line, _ int) Line {
		item.LineNumber = next
		next += 1
		return item
	})

	return nil
}

func (p *ParsedContent) ParseContent(content string) error {
	err := p.importContent(content)
	if err != nil {
		return err
	}

	err = p.categorizeLines()
	if err != nil {
		return err
	}

	err = p.compactLines()
	if err != nil {
		return err
	}

	return nil
}

func (p *ParsedContent) TransposeUpOneStep() {
	for lineIndex := range p.Lines {
		if p.Lines[lineIndex].Type == LineTypes.CHORDS {
			for partIndex := range p.Lines[lineIndex].Parts {
				if p.Lines[lineIndex].Parts[partIndex].Type == LetterRunTypes.SEPARATORRUN {
					p.Lines[lineIndex].Parts[partIndex].Letters = p.Lines[lineIndex].Parts[partIndex].OriginalLetters
				}
			}

			longer := make([]int, 0)
			shorter := make([]int, 0)
			for partIndex := range p.Lines[lineIndex].Parts {
				if p.Lines[lineIndex].Parts[partIndex].Type == LetterRunTypes.CHORDRUN {
					p.Lines[lineIndex].Parts[partIndex].Chord.StepUp()
					newLetters := p.Lines[lineIndex].Parts[partIndex].Chord.String()
					p.Lines[lineIndex].Parts[partIndex].TransposedLetters = newLetters

					if len(newLetters) > len(p.Lines[lineIndex].Parts[partIndex].Chord.OriginalString) {
						longer = append(longer, partIndex)
					} else {
						if len(newLetters) < len(p.Lines[lineIndex].Parts[partIndex].Chord.OriginalString) {
							shorter = append(shorter, partIndex)
						}
					}
				}
			}

			for _, index := range longer {
				if index < len(p.Lines[lineIndex].Parts)-1 {
					if p.Lines[lineIndex].Parts[index+1].Type == LetterRunTypes.SEPARATORRUN {
						text := p.Lines[lineIndex].Parts[index+1].Letters
						if len(text) > 1 && text[0] == ' ' && text[1] == ' ' {
							p.Lines[lineIndex].Parts[index+1].Letters = p.Lines[lineIndex].Parts[index+1].Letters[1:]
						}
					}
				}
			}

			for _, index := range shorter {
				if index < len(p.Lines[lineIndex].Parts)-1 {
					p.Lines[lineIndex].Parts[index].TransposedLetters = p.Lines[lineIndex].Parts[index].TransposedLetters + " "
				}
			}
		}
	}
}

func (p *ParsedContent) TransposeDownOneStep() {
	for lineIndex := range p.Lines {
		if p.Lines[lineIndex].Type == LineTypes.CHORDS {
			for partIndex := range p.Lines[lineIndex].Parts {
				if p.Lines[lineIndex].Parts[partIndex].Type == LetterRunTypes.SEPARATORRUN {
					p.Lines[lineIndex].Parts[partIndex].Letters = p.Lines[lineIndex].Parts[partIndex].OriginalLetters
				}
			}

			longer := make([]int, 0)
			shorter := make([]int, 0)
			for partIndex := range p.Lines[lineIndex].Parts {
				if p.Lines[lineIndex].Parts[partIndex].Type == LetterRunTypes.CHORDRUN {
					p.Lines[lineIndex].Parts[partIndex].Chord.StepDown()
					newLetters := p.Lines[lineIndex].Parts[partIndex].Chord.String()
					p.Lines[lineIndex].Parts[partIndex].TransposedLetters = newLetters

					if len(newLetters) > len(p.Lines[lineIndex].Parts[partIndex].Chord.OriginalString) {
						longer = append(longer, partIndex)
					} else {
						if len(newLetters) < len(p.Lines[lineIndex].Parts[partIndex].Chord.OriginalString) {
							shorter = append(shorter, partIndex)
						}
					}
				}
			}

			for _, index := range longer {
				if index < len(p.Lines[lineIndex].Parts)-1 {
					if p.Lines[lineIndex].Parts[index+1].Type == LetterRunTypes.SEPARATORRUN {
						text := p.Lines[lineIndex].Parts[index+1].Letters
						if len(text) > 1 && text[0] == ' ' && text[1] == ' ' {
							p.Lines[lineIndex].Parts[index+1].Letters = p.Lines[lineIndex].Parts[index+1].Letters[1:]
						}
					}
				}
			}

			for _, index := range shorter {
				if index < len(p.Lines[lineIndex].Parts)-1 {
					p.Lines[lineIndex].Parts[index].TransposedLetters = p.Lines[lineIndex].Parts[index].TransposedLetters + " "
				}
			}
		}
	}
}
