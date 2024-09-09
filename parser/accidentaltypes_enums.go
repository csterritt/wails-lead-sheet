// Code generated by goenums. DO NOT EDIT.
// This file was generated by github.com/zarldev/goenums
// using the command:
// goenums accidental-type.go

package parser

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
)

type AccidentalType struct {
	accidentalType
}

type accidentaltypesContainer struct {
	NONE  AccidentalType
	SHARP AccidentalType
	FLAT  AccidentalType
}

var AccidentalTypes = accidentaltypesContainer{
	NONE: AccidentalType{
		accidentalType: None,
	},
	SHARP: AccidentalType{
		accidentalType: Sharp,
	},
	FLAT: AccidentalType{
		accidentalType: Flat,
	},
}

func (c accidentaltypesContainer) All() []AccidentalType {
	return []AccidentalType{
		c.NONE,
		c.SHARP,
		c.FLAT,
	}
}

var invalidAccidentalType = AccidentalType{}

func ParseAccidentalType(a any) (AccidentalType, error) {
	res := invalidAccidentalType
	switch v := a.(type) {
	case AccidentalType:
		return v, nil
	case []byte:
		res = stringToAccidentalType(string(v))
	case string:
		res = stringToAccidentalType(v)
	case fmt.Stringer:
		res = stringToAccidentalType(v.String())
	case int:
		res = intToAccidentalType(v)
	case int64:
		res = intToAccidentalType(int(v))
	case int32:
		res = intToAccidentalType(int(v))
	}
	return res, nil
}

func stringToAccidentalType(s string) AccidentalType {
	switch s {
	case "None":
		return AccidentalTypes.NONE
	case "Sharp":
		return AccidentalTypes.SHARP
	case "Flat":
		return AccidentalTypes.FLAT
	}
	return invalidAccidentalType
}

func intToAccidentalType(i int) AccidentalType {
	if i < 0 || i >= len(AccidentalTypes.All()) {
		return invalidAccidentalType
	}
	return AccidentalTypes.All()[i]
}

func ExhaustiveAccidentalTypes(f func(AccidentalType)) {
	for _, p := range AccidentalTypes.All() {
		f(p)
	}
}

var validAccidentalTypes = map[AccidentalType]bool{
	AccidentalTypes.NONE:  true,
	AccidentalTypes.SHARP: true,
	AccidentalTypes.FLAT:  true,
}

func (p AccidentalType) IsValid() bool {
	return validAccidentalTypes[p]
}

func (p AccidentalType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

func (p *AccidentalType) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(bytes.Trim(b, `"`), ` `)
	newp, err := ParseAccidentalType(b)
	if err != nil {
		return err
	}
	*p = newp
	return nil
}

func (p *AccidentalType) Scan(value any) error {
	newp, err := ParseAccidentalType(value)
	if err != nil {
		return err
	}
	*p = newp
	return nil
}

func (p AccidentalType) Value() (driver.Value, error) {
	return p.String(), nil
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the goenums command to generate them again.
	// Does not identify newly added constant values unless order changes
	var x [1]struct{}
	_ = x[None-0]
	_ = x[Sharp-1]
	_ = x[Flat-2]
}

const _accidentaltypes_name = "NoneSharpFlat"

var _accidentaltypes_index = [...]uint16{0, 4, 9, 13}

func (i accidentalType) String() string {
	if i < 0 || i >= accidentalType(len(_accidentaltypes_index)-1) {
		return "accidentaltypes(" + (strconv.FormatInt(int64(i), 10) + ")")
	}
	return _accidentaltypes_name[_accidentaltypes_index[i]:_accidentaltypes_index[i+1]]
}
