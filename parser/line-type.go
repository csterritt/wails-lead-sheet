package parser

//go:generate goenums line-type.go

type lineType int

const (
	Text lineType = iota
	Section
	Chords
	Lyrics
	Empty
)
