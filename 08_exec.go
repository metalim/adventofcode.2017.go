package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

// 1   2     4  5  6

var test2 = ``

type regs map[string]int
type comp func(rs regs, r string, v int) bool
type op func(rs regs, r string, v int)

var comps = map[string]comp{
	"==": func(rs regs, r string, v int) bool { return rs[r] == v },
	"!=": func(rs regs, r string, v int) bool { return rs[r] != v },
	"<=": func(rs regs, r string, v int) bool { return rs[r] <= v },
	">=": func(rs regs, r string, v int) bool { return rs[r] >= v },
	"<":  func(rs regs, r string, v int) bool { return rs[r] < v },
	">":  func(rs regs, r string, v int) bool { return rs[r] > v },
}

var ops = map[string]op{
	"inc": func(rs regs, r string, v int) { rs[r] += v },
	"dec": func(rs regs, r string, v int) { rs[r] -= v },
}

func main() {
	for p := range source.Test(1, test1, `1`).Test(2, test1, `10`).Advent(2017, 8) {
		fmt.Println(Brown("\n" + p.Name))
		ssw := p.Lines().Words()
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		var maxmax int
		mr := regs{}
		for _, l := range ssw {
			val, _ := strconv.Atoi(l[2])
			cval, _ := strconv.Atoi(l[6])
			if comps[l[5]](mr, l[4], cval) {
				ops[l[1]](mr, l[0], val)
				if maxmax < mr[l[0]] {
					maxmax = mr[l[0]]
				}
			}
		}

		if p.Part(1) {

			var max int
			for _, val := range mr {
				if max < val {
					max = val
				}
			}
			p.SubmitInt(1, max)
		}

		if p.Part(2) {
			p.SubmitInt(2, maxmax)
		}
	}
}
