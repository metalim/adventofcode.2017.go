package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0
3
0
1
-3`

func main() {
	for p := range source.Test(1, test1, `5`).Test(2, test1, `10`).Advent(2017, 5) {
		fmt.Println(Brown("\n" + p.Name))

		if p.Part(1) {
			sn := p.Ints()
			var i, steps int
			for i >= 0 && i < len(sn) {
				steps++
				j := i + sn[i]
				sn[i]++
				i = j
			}
			p.SubmitInt(1, steps)
		}
		if p.Part(2) {
			sn := p.Ints()
			var i, steps int

			for i >= 0 && i < len(sn) {
				steps++
				j := i + sn[i]
				if sn[i] >= 3 {
					sn[i]--
				} else {
					sn[i]++
				}
				i = j
			}
			p.SubmitInt(2, steps)
		}
	}
}
