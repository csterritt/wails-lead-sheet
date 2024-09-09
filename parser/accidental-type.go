package parser

//go:generate goenums accidental-type.go

type accidentalType int

const (
	None accidentalType = iota
	Sharp
	Flat
)
