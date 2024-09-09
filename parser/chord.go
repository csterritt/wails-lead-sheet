package parser

import "strings"

type Chord struct {
	Note       string
	Accidental AccidentalType
	Flavor     string
}

func (c Chord) String() string {
	res := c.Note

	if c.Accidental == AccidentalTypes.SHARP {
		res += "#"
	}

	if c.Accidental == AccidentalTypes.FLAT {
		res += "b"
	}

	res += c.Flavor

	return res
}

func MakeChord(original string) Chord {
	res := Chord{}
	original = strings.ToLower(original)
	if original[0] < 'a' || original[0] > 'g' {
		return res
	}

	if len(original) == 1 {
		res.Note = strings.ToUpper(original)
		return res
	}

	start := 1
	if original[1] == 'b' || original[1] == '#' {
		start = 2
	}

	res.Note = strings.ToUpper(original[:1])
	if original[1] == '#' {
		res.Accidental = AccidentalTypes.SHARP
	}

	if original[1] == 'b' {
		res.Accidental = AccidentalTypes.FLAT
	}

	if len(original) == 2 {
		return res
	}

	_, found := knownChordSuffixes[original[start:]]
	if found {
		res.Flavor = original[start:]
	}

	return res
}
