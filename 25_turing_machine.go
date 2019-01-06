package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"metalim/advent/2017/lib/turing"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `3`)

	for par := range ins.Advent(2017, 25) {
		fmt.Println(Brown("\n" + par.Name))
		ssw := par.Lines().WordsTrim(":.")
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		state := ssw[0][len(ssw[0])-1]
		n, _ := strconv.Atoi(ssw[1][len(ssw[1])-2])
		ssw = ssw[2:]

		rules := map[string][2]cmd{}
		for len(ssw) > 0 {
			ins := ssw[1][len(ssw[1])-1]

			// assuming ssw[2] is 0.
			v1, _ := strconv.Atoi(ssw[3][len(ssw[3])-1])
			l1 := ssw[4][len(ssw[4])-1] == "left"
			s1 := ssw[5][len(ssw[5])-1]

			// assuming ssw[6] is 1.
			v2, _ := strconv.Atoi(ssw[7][len(ssw[7])-1])
			l2 := ssw[8][len(ssw[8])-1] == "left"
			s2 := ssw[9][len(ssw[9])-1]

			rules[ins] = [2]cmd{{v1, l1, s1}, {v2, l2, s2}}
			ssw = ssw[10:]
		}

		t := turing.Tape{}

		// execute.
		for ; n > 0; n-- {
			c := rules[state][t.Get()]
			t.Set(c.v)
			if c.l {
				t.Left()
			} else {
				t.Right()
			}
			state = c.s
		}

		// count checksum.
		var sum int
		l, r := t.Bounds()
		for ; l <= r; l++ {
			if t.GetAt(l) == 1 {
				sum++
			}
		}
		par.SubmitInt(1, sum)

		if par.Part(2) {
			par.SubmitInt(2, 0)
		}
	}
}

////////////////////////////////////////////////////////////////////////

type cmd struct {
	v int
	l bool
	s string
}
