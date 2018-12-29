package main

import (
	"fmt"
	"metalim/advent/2017/lib/debug"
	"metalim/advent/2017/lib/source"
	"time"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0: 3
1: 2
4: 4
6: 4`
var test2 = ``

type layer struct {
	x, y, d, l int
}

func firewall(ssn [][]int) (map[int]*layer, int) {
	fw := map[int]*layer{}
	var max int
	for _, l := range ssn {
		fw[l[0]] = &layer{x: l[0], l: l[1], d: 1}
		max = l[0]
	}
	return fw, max
}

func main() {
	// source.Dry()

	var ins source.Inputs

	ins = ins.Test(1, test1, `24`)
	ins = ins.Test(2, test1, `10`)

	for p := range ins.Advent(2017, 13) {

		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()
		fmt.Println(len(ssn), Black(ssn[:3]).Bold())

		if p.Part(1) {
			var sev int
			fw, max := firewall(ssn)

			for x := 0; x <= max; x++ {
				if s, ok := fw[x]; ok {
					if s.y == 0 { // cought?
						sev += s.x * s.l
					}
				}
				for _, s := range fw { // move scanners
					if s.y+s.d < 0 || s.y+s.d >= s.l {
						s.d = -s.d
					}
					s.y += s.d
				}
			}
			p.SubmitInt(1, sev)
		}

		if p.Part(2) {
			fw, max := firewall(ssn)
			var x0 int

		NEXT:
			for x0 = 0; ; x0++ {
				debug.LogT(time.Second, Black(x0).Bold())

				// pre-move scanners.
				for _, s := range fw {
					mod := (s.l - 1) * 2
					s.y = x0 % mod
					s.d = 1
					if s.y >= s.l {
						s.y = mod - s.y
						s.d = -1
					}
				}

				for x := 0; x <= max; x++ {
					if x >= 0 {
						if s, ok := fw[x]; ok {
							if s.y == 0 { // cought?
								continue NEXT
							}
						}
					}
					for _, s := range fw { // move scanners
						if s.y+s.d < 0 || s.y+s.d >= s.l {
							s.d = -s.d
						}
						s.y += s.d
					}
				}
				break // found it.
			}
			p.SubmitInt(2, x0)
		}

	}
}
