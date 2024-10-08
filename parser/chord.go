package parser

import "strings"

type Chord struct {
	Note               string
	BassNote           *Chord
	Accidental         AccidentalType
	Flavor             string
	OriginalString     string
	OriginalAccidental AccidentalType
}

func nextUp(note string) string {
	res := "X"
	switch note {
	case "A":
		res = "B"
	case "C":
		res = "D"
	case "D":
		res = "E"
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
	case "D":
		res = "C"
	case "E":
		res = "D"
	case "G":
		res = "F"
	}

	return res
}

func (c *Chord) String() string {
	res := c.Note

	if c.Accidental == AccidentalTypes.SHARP {
		res += "#"
	}

	if c.Accidental == AccidentalTypes.FLAT {
		res += "b"
	}

	res += c.Flavor

	if c.BassNote != nil {
		res += "/" + c.BassNote.String()
	}

	return res
}

func MakeChord(original string) Chord {
	res := Chord{}
	if len(original) == 0 {
		return res
	}

	copyOfOriginal := strings.ToLower(original)
	if copyOfOriginal == "n.c." {
		return res
	}

	if strings.Count(copyOfOriginal, "/") > 1 {
		return res
	}

	bassNote := ""
	spot := strings.Index(copyOfOriginal, "/")
	if spot != -1 {
		bassNote = original[spot+1:]
		copyOfOriginal = copyOfOriginal[:spot]
	}

	if copyOfOriginal[0] < 'a' || copyOfOriginal[0] > 'g' {
		return res
	}

	if len(copyOfOriginal) == 1 {
		res.Note = strings.ToUpper(original[:1])
		if bassNote != "" {
			bassNoteChord := MakeChord(bassNote)
			res.BassNote = &bassNoteChord
		}

		res.OriginalString = res.String()
		res.OriginalAccidental = res.Accidental

		return res
	}

	start := 1
	hasAccidental := false
	if copyOfOriginal[1] == 'b' || copyOfOriginal[1] == '#' {
		start = 2
		hasAccidental = true
	}

	res.Note = strings.ToUpper(original[:1])
	if copyOfOriginal[1] == '#' {
		res.Accidental = AccidentalTypes.SHARP
	}

	if copyOfOriginal[1] == 'b' {
		res.Accidental = AccidentalTypes.FLAT
	}

	if hasAccidental && len(copyOfOriginal) == 2 {
		if bassNote != "" {
			bassNoteChord := MakeChord(bassNote)
			res.BassNote = &bassNoteChord
		}

		res.OriginalString = res.String()
		res.OriginalAccidental = res.Accidental

		return res
	}

	_, found := knownChordSuffixes[copyOfOriginal[start:]]
	if found {
		res.Flavor = copyOfOriginal[start:]
	} else {
		return Chord{}
	}

	if bassNote != "" {
		bassNoteChord := MakeChord(bassNote)
		res.BassNote = &bassNoteChord
	}

	res.OriginalString = res.String()
	res.OriginalAccidental = res.Accidental

	return res
}

func (c *Chord) Reset() {
	c.Note = c.OriginalString
	c.Accidental = c.OriginalAccidental
}

func (c *Chord) StepUp() {
	if c.BassNote != nil {
		c.BassNote.StepUp()
	}

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
	if c.BassNote != nil {
		c.BassNote.StepDown()
	}

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
