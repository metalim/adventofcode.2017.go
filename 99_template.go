package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = ``
var test2 = ``

func main() {
	source.Dry()

	var ins source.Inputs

	ins = ins.Test(1, test1, `11`)
	ins = ins.Test(2, test2, `22`)

	for p := range ins.Advent(2017, 99) {
		fmt.Println(Brown("\n" + p.Name))
		ssw := p.Lines().Words()
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		if p.Part(1) {
			p.SubmitInt(1, 1)
		}

		// if p.Part(2) {
		// 	p.SubmitInt(2, 2)
		// }
	}
}
