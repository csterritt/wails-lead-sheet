package parser

//go:generate goenums accidental-type.go

type accidentalType int

const (
	Natural accidentalType = iota
	Sharp
	Flat
)
