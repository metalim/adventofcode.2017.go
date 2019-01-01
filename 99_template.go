package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = ``

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `1`)
	// ins = ins.Test(2, test2, `2`)

	for par := range ins.Advent(2017, 99) {
		fmt.Println(Brown("\n" + par.Name))
		ssw := par.Lines().Words()
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		if par.Part(1) {
			par.DrySubmitInt(1, 1)
		}

		// if par.Part(2) {
		// 	par.DrySubmitInt(2, 2)
		// }
	}
}
