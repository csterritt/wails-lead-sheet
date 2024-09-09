package parser

import "strings"

type Chord struct {
	Note       string
	Accidental AccidentalType
	Flavor     string
}

func nextUp(note string) string {
	res := "X"
	switch note {
	case "A":
		res = "B"
	case "B":
		res = "C"
	case "C":
		res = "D"
	case "D":
		res = "E"
	case "E":
		res = "F"
	case "F":
		res = "G"
	case "G":
		res = "A"
	}

	return res
}

func nextDown(note string) string {
	res := "X"
	switch note {
	case "A":
		res = "G"
	case "B":
		res = "A"
	case "C":
		res = "B"
	case "D":
		res = "C"
	case "E":
		res = "D"
	case "F":
		res = "E"
	case "G":
		res = "F"
	}

	return res
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
	hasAccidental := false
	if original[1] == 'b' || original[1] == '#' {
		start = 2
		hasAccidental = true
	}

	res.Note = strings.ToUpper(original[:1])
	if original[1] == '#' {
		res.Accidental = AccidentalTypes.SHARP
	}

	if original[1] == 'b' {
		res.Accidental = AccidentalTypes.FLAT
	}

	if hasAccidental && len(original) == 2 {
		return res
	}

	_, found := knownChordSuffixes[original[start:]]
	if found {
		res.Flavor = original[start:]
	}

	return res
}

func (c *Chord) StepUp() {
	if c.Note == "B" && c.Accidental == AccidentalTypes.NATURAL {
		c.Note = "C"
		return
	}

	if c.Note == "E" && c.Accidental == AccidentalTypes.NATURAL {
		c.Note = "F"
		return
	}

	switch c.Accidental {
	case AccidentalTypes.NATURAL:
		c.Accidental = AccidentalTypes.SHARP
	case AccidentalTypes.FLAT:
		c.Accidental = AccidentalTypes.NATURAL
	case AccidentalTypes.SHARP:
		c.Accidental = AccidentalTypes.NATURAL
		c.Note = nextUp(c.Note)
	}
}

func (c *Chord) StepDown() {
	if c.Note == "C" && c.Accidental == AccidentalTypes.NATURAL {
		c.Note = "B"
		return
	}

	if c.Note == "F" && c.Accidental == AccidentalTypes.NATURAL {
		c.Note = "E"
		return
	}

	switch c.Accidental {
	case AccidentalTypes.NATURAL:
		c.Accidental = AccidentalTypes.FLAT
	case AccidentalTypes.SHARP:
		c.Accidental = AccidentalTypes.NATURAL
	case AccidentalTypes.FLAT:
		c.Accidental = AccidentalTypes.NATURAL
		c.Note = nextDown(c.Note)
	}
}
