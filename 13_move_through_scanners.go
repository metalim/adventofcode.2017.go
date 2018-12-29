package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0: 3
1: 2
4: 4
6: 4`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `24`)
	ins = ins.Test(2, test1, `10`)

	for p := range ins.Advent(2017, 13) {

		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()
		fmt.Println(len(ssn), Black(ssn[:3]).Bold())

		type layer struct {
			x, l int
		}
		fw := map[int]layer{}
		for _, l := range ssn {
			fw[l[0]] = layer{x: l[0], l: l[1]}
		}

		if p.Part(1) {
			var sev int

			// instead of moving scanners, calculate their position when packet is passing.
			for _, s := range fw {
				if s.x%(2*s.l-2) == 0 {
					sev += s.x * s.l
				}
			}
			p.SubmitInt(1, sev)
		}

		if p.Part(2) {
			var x0 int

		DELAY:
			for x0 = 0; ; x0++ { // try delays, until solution is found.

				// instead of moving scanners, calculate their position when packet is passing.
				for _, s := range fw {
					if (x0+s.x)%(2*s.l-2) == 0 {
						continue DELAY
					}
				}
				break // found it
			}
			p.SubmitInt(2, x0)
		}

	}
}
