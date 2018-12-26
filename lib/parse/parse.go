package parse

import "metalim/advent/2017/lib/advent"

// Parser chain.
type Parser struct {
	name, val string
}

// ParserNum chain.
type ParserNum Parser

// ParserSplit chain.
type ParserSplit Parser

////////////////////////////////////////////////////////////////////////
// Sources
//

// Inputs gets input files from Advent of Code website.
func Inputs(year, day int) <-chan Parser {
	ch := make(chan Parser)
	go func() {

		for in := range advent.Inputs(year, day) {
			ch <- Parser{name: in.Name, val: in.Val}
		}

		close(ch)
	}()
	return ch
}

////////////////////////////////////////////////////////////////////////
// Operations
//

// Trim space.
func (p Parser) Trim() Parser {
	return p
}

// Split by separator.
func (p Parser) Split(sep string) ParserSplit {
	return ParserSplit(p)
}

// Numbers from string.
func (p Parser) Numbers() ParserNum {
	return ParserNum(p)
}

// Value .
func (p Parser) Value() string {
	return p.val
}
