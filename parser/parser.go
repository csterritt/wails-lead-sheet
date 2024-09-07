package parser

import (
	"strings"

	"github.com/samber/lo"
)

type ParsedContent struct {
	Lines []string
}

func (p *ParsedContent) ParseContent(content string) error {
	p.Lines = lo.FilterMap(strings.Split(content, "\n"), func(s string, _ int) (string, bool) {
		res := strings.TrimRight(s, " \t\r\n")
		return res, len(res) > 0
	})

	return nil
}
