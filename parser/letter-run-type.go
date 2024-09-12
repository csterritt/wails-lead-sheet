package parser

//go:generate goenums letter-run-type.go

type letterRunType int

const (
	WordRun letterRunType = iota
	ChordRun
	SeparatorRun
	UnknownRun
)
