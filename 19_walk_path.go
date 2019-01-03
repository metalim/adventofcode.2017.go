package main

import (
	"fmt"
	"metalim/advent/2017/lib/field"
	"metalim/advent/2017/lib/source"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var test1 = `.     |          .
.     |  +--+    .
.     A  |  C    .
. F---|----E|--+ .
.     |  |  |  D .
.     +B-+  +--+ .
`

func main() {
	var ins source.Inputs

	ins = ins.Test(3, test1, `ABCDEF`, `38`)

	for par := range ins.Advent(2017, 19) {
		fmt.Println(Brown("\n" + par.Name))
		var f field.Slice
		f.SetDefault(' ')
		field.FillFromString(&f, field.Pos{}, par.Val)
		startX := strings.IndexRune(par.Val, '|')
		fmt.Printf(Black("%v, startX: %d\n").Bold().String(), f.Bounds(), startX)

		var out []byte

		start := field.Pos{startX, 0}

		canStepOn := func(p field.Pos) bool { return f.Get(p) != ' ' }
		stepOn := func(p field.Pos, d field.Dir8) int {
			c := byte(f.Get(p))
			if strings.IndexByte("|-+", c) == -1 {
				out = append(out, byte(c))
			}

			if f.Get(p) == '+' {
				return 1<<((d+2)&7) + 1<<((d-2)&7) // turn right or left
			}
			return 1 << d // continue in same direction
		}
		_, steps := field.Walk(start, field.Dir80P, canStepOn, stepOn)

		if par.Part(1) {
			par.Submit(1, string(out))
		}

		if par.Part(2) {
			par.SubmitInt(2, steps)
		}

	}
}
