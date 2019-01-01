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
		f.FillFromString(field.Pos{}, par.Val)
		start := strings.IndexRune(par.Val, '|')
		fmt.Printf(Black("%v, start: %d\n").Bold().String(), f.Bounds(), start)

		var out []byte

		p := field.Pos{start, 0}
		d := 1
		var step int
		for {
			step++
			c := f.Get(p)
			switch c {
			case '|', '-', '+':
			default:
				out = append(out, byte(c))
			}
			p1 := field.Step(p, d)
			if f.Get(p1) != ' ' {
				p = p1
				continue
			}
			if c == '+' {
				dr := (d + 1) % 4
				pr := field.Step(p, dr)
				if f.Get(pr) != ' ' {
					d = dr
					p = pr
					continue
				}
				dl := (d + 3) % 4
				pl := field.Step(p, dl)
				if f.Get(pl) != ' ' {
					d = dl
					p = pl
					continue
				}
			}
			break
		}

		if par.Part(1) {
			par.Submit(1, string(out))
		}

		if par.Part(2) {
			par.SubmitInt(2, step)
		}
	}
}
