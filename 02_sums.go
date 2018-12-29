package main

import (
	"fmt"
	"metalim/advent/2017/lib/numbers"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `5 1 9 5
7 5 3
2 4 6 8`

var test2 = `5 9 2 8
9 4 7 3
3 8 6 5`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `18`)
	ins = ins.Test(2, test2, `9`)

	for p := range ins.Advent(2017, 2) {
		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()

		if p.Part(1) {
			var sum int
			for _, l := range ssn {
				min, max := numbers.MinMax(l)
				sum += max - min
			}
			p.SubmitInt(1, sum)
		}

		if p.Part(2) {
			var sum int
			for _, l := range ssn {
				for i, a := range l {
					for _, b := range l[i+1:] {
						if a%b == 0 {
							sum += a / b
						} else if b%a == 0 {
							sum += b / a
						}
					}
				}
			}
			p.SubmitInt(2, sum)
		}
	}
}
