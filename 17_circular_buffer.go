package main

import (
	"fmt"
	"metalim/advent/2017/lib/circular"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `3`, `638`)

	for p := range ins.Advent(2017, 17) {
		fmt.Println(Brown("\n" + p.Name))
		step := p.Ints()[0]
		fmt.Println(Black(step).Bold())

		if p.Part(1) {
			n := circular.NewList(0) // circular.NewSlice(0) works too, but is slower for large number of nodes.
			for i := 1; i <= 2017; i++ {
				n.Skip(step)
				n.InsertAfter(i)
			}
			p.SubmitInt(1, n.Skip(1).Get())
		}

		if p.Part(2) {
			var after0, pos0, pos int
			length := 1
			for i := 1; i <= 5e7; i++ {
				pos = (pos + step) % length
				if pos < pos0 {
					pos0++
				} else if pos == pos0 {
					after0 = i
				}
				pos++
				length++
			}
			p.SubmitInt(2, after0)
		}
	}
}
