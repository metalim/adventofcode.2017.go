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
	for p := range source.Test(1, test1, `11`).Test(2, test2, `22`).Advent(2017, 99) {
		fmt.Println(Brown("\n" + p.Name))
		ssn := p.Lines().Ints()
		fmt.Println(Black(ssn).Bold())

		if p.Part(1) {
			p.SubmitInt(1, 1)
		}

		// if p.Part(2) {
		// 	p.SubmitInt(2, 2)
		// }
	}
}
