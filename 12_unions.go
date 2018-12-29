package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"metalim/advent/2017/lib/union"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `6`)
	ins = ins.Test(2, test1, `2`)

	for p := range ins.Advent(2017, 12) {
		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()
		fmt.Println(len(ssn), Black(ssn[0:3]).Bold())

		u := union.New()
		for _, sn := range ssn {
			u.Link(sn...)
		}

		if p.Part(1) {
			p.SubmitInt(1, len(u.Unions[u.Nodes[0]]))
		}

		if p.Part(2) {
			p.SubmitInt(2, len(u.Unions))
		}

	}
}
