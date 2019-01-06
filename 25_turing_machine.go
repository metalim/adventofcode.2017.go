package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
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

			v1, _ := strconv.Atoi(ssw[3][len(ssw[3])-1])
			l1 := ssw[4][len(ssw[4])-1] == "left"
			st1 := ssw[5][len(ssw[5])-1]

			v2, _ := strconv.Atoi(ssw[7][len(ssw[7])-1])
			l2 := ssw[8][len(ssw[8])-1] == "left"
			st2 := ssw[9][len(ssw[9])-1]

			rules[ins] = [2]cmd{{v1, l1, st1}, {v2, l2, st2}}
			ssw = ssw[10:]
		}

		tape := tape{}

		for ; n > 0; n-- {
			c := rules[state][tape.get()]
			tape.set(c.v)
			if c.l {
				tape.left()
			} else {
				tape.right()
			}
			state = c.st
		}

		var sum int
		l, r := tape.bounds()
		for ; l <= r; l++ {
			if tape.getAt(l) == 1 {
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
	v  int
	l  bool
	st string
}

////////////////////////////////////////////////////////////////////////

type tape struct {
	d [2][]int
	p int
}

func (t *tape) sel(p int) (int, int) {
	if p < 0 {
		return 1, -p - 1
	}
	return 0, p
}

// getAt - for iteration.
func (t *tape) getAt(pos int) int {
	i, p := t.sel(pos)
	if p < len(t.d[i]) {
		return t.d[i][p]
	}
	return 0
}

func (t *tape) get() int {
	return t.getAt(t.p)
}

func (t *tape) set(v int) {
	i, p := t.sel(t.p)
	if p >= len(t.d[i]) {
		t.d[i] = append(t.d[i], make([]int, p+1-len(t.d[i]))...) // extra space
	}
	t.d[i][p] = v
}

func (t *tape) left() {
	t.p--
}

func (t *tape) right() {
	t.p++
}

func (t *tape) go2(p int) {
	t.p = p
}

func (t *tape) goRight(d int) {
	t.p += d
}

// bounds inclusive
func (t *tape) bounds() (int, int) {
	return -len(t.d[1]), len(t.d[0]) - 1
}
