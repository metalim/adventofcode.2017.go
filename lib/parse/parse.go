package parse

import (
	"metalim/advent/2017/lib/advent"
	"regexp"
	"strconv"
	"strings"
)

// Parser chain.
type Parser struct {
	Parts int
	advent.Input
}

// ParserNum chain.
type ParserNum Parser

// ParserSplit chain.
type ParserSplit struct {
	Parser
	Values []string
}

// Test chain.
type Test struct {
	parts int
	string
}

// Prepend chain of Tests.
type Prepend struct {
	tests []Test
}

////////////////////////////////////////////////////////////////////////
// Sources
//

// Advent gets input files from Advent of Code website.
func Advent(year, day int) <-chan Parser {
	ch := make(chan Parser)
	go func() {

		for in := range advent.Inputs(year, day) {
			ch <- Parser{Parts: 1 + 2, Input: in}
		}

		close(ch)
	}()
	return ch
}

// Tests prepends test strings to the inputs.
func Tests(parts int, ss ...string) Prepend {
	st := make([]Test, 0, len(ss))
	for _, s := range ss {
		st = append(st, Test{parts, s})
	}
	return Prepend{st}
}

// Tests prepends test strings to the chain.
func (p Prepend) Tests(parts int, ss ...string) Prepend {
	for _, s := range ss {
		p.tests = append(p.tests, Test{parts, s})
	}
	return p
}

// Advent with Test prepended, gets input files from Advent of Code website.
func (p Prepend) Advent(year, day int) <-chan Parser {
	ch := make(chan Parser)

	go func() {

		for i, t := range p.tests {
			ch <- Parser{Parts: t.parts, Input: advent.Input{Name: "test" + strconv.Itoa(i), Val: t.string}}
		}

		for in := range advent.Inputs(year, day) {
			ch <- Parser{Parts: 1 + 2, Input: in}
		}

		close(ch)
	}()
	return ch
}

////////////////////////////////////////////////////////////////////////
// Operations
//

// Part filter.
func (p Parser) Part(part int) bool {
	return p.Parts&part != 0
}

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
