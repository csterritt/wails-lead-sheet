package parser

import (
	"strings"

	"github.com/samber/lo"
)

type Line struct {
	Text string
	Type LineType
}

type ParsedContent struct {
	Lines []Line
}

func (p *ParsedContent) importContent(content string) error {
	p.Lines = lo.FilterMap(strings.Split(content, "\n"), func(s string, _ int) (Line, bool) {
		res := strings.TrimRight(s, " \t\r\n")
		return Line{Text: res, Type: LineTypes.TEXT}, len(res) > 0
	})

	return nil
}
