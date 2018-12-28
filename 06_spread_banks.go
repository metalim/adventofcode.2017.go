package main

import (
	"fmt"
	"metalim/advent/2017/lib/numbers"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0 2 7 0`

func main() {
	// source.Dry()
	for p := range source.Test(1, test1, `5`).Test(2, test1, `4`).Advent(2017, 6) {
		fmt.Println(Brown("\n" + p.Name))
		sn := p.Ints()
		fmt.Println(Black(sn).Bold())

		seen := map[string]int{}
		var step int
		for seen[numbers.Join(sn, ":")] == 0 {
			seen[numbers.Join(sn, ":")] = step
			step++
			_, max := numbers.MinMax(sn)
			var i, n int
			for i, n = range sn {
				if n == max {
					break
				}
			}
			sn[i] = 0
			for ; max > 0; max-- {
				i = (i + 1) % len(sn)
				sn[i]++
			}
		}

		if p.Part(1) {
			p.SubmitInt(1, step)
		}

		if p.Part(2) {
			p.SubmitInt(2, step-seen[numbers.Join(sn, ":")])
		}
	}
}
