package source

import (
	"regexp"
	"strconv"
	"strings"
)

// Parser chain.
type Parser struct {
	Input
}

// ParserSplit chain.
type ParserSplit struct {
	Parser
	Values []string
}

////////////////////////////////////////////////////////////////////////
// Operations
//

// Ints from string.
func (p Parser) Ints() []int {
	ss := rInts.FindAllString(p.Val, -1)
	sn := make([]int, len(ss))
	for i, s := range ss {
		sn[i], _ = strconv.Atoi(s)
	}
	return sn
}

// Split by separator.
func (p Parser) Split(sep string) ParserSplit {
	ss := strings.Split(p.Input.Val, sep)
	return ParserSplit{Parser: p, Values: ss}
}

// Lines split by "\n".
func (p Parser) Lines() ParserSplit {
	return p.Split("\n")
}

var rInts = regexp.MustCompile("-?\\d+")

// Ints from []string.
func (p ParserSplit) Ints() [][]int {
	out := make([][]int, 0, len(p.Values))
	for _, l := range p.Values {
		ss := rInts.FindAllString(l, -1)
		sn := make([]int, len(ss))
		for i, s := range ss {
			sn[i], _ = strconv.Atoi(s)
		}
		out = append(out, sn)
	}
	return out
}
