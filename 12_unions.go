package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

var test2 = ``

func main() {
	// source.Dry()
	ins := source.Test(1, test1, `6`)
	ins = ins.Test(2, test1, `2`)
	for p := range ins.Advent(2017, 12) {
		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()
		fmt.Println(len(ssn), Black(ssn[0:3]).Bold())

		if p.Part(1) {
			q := []int{0}
			seen := map[int]bool{0: true}
			var count int
			for len(q) > 0 {
				n := q[0]
				q = q[1:]
				count++
				for _, n2 := range ssn[n][1:] {
					if !seen[n2] {
						seen[n2] = true
						q = append(q, n2)
					}
				}
			}
			p.SubmitInt(1, count)
		}

		if p.Part(2) {
			seen := map[int]bool{}
			var groups int
			for i := range ssn {
				if seen[i] {
					continue
				}
				groups++
				q := []int{i}

				for len(q) > 0 {
					n := q[0]
					q = q[1:]
					for _, n2 := range ssn[n][1:] {
						if !seen[n2] {
							seen[n2] = true
							q = append(q, n2)
						}
					}
				}
			}
			p.SubmitInt(2, groups)

		}
	}
}
