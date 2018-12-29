package main

import (
	"fmt"
	"metalim/advent/2017/lib/numbers"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var abs = numbers.Abs

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `ne,ne,sw,sw`, `0`)
	ins = ins.Test(1, `se,sw,se,sw,sw`, `3`)

	for par := range ins.Advent(2017, 11) {
		fmt.Println(Brown("\n" + par.Name))
		sw := par.Lines().Split(",").Values
		fmt.Println(len(sw), Black(sw[:4]).Bold())

		var x, y, z, d, max int
		for _, dir := range sw {
			switch dir {
			case "n":
				y++
			case "s":
				y--
			case "ne":
				x++
			case "sw":
				x--
			case "nw":
				x--
				y++
			case "se":
				x++
				y--
			}
			z = -x - y
			d = (abs(x) + abs(y) + abs(z)) / 2
			if max < d {
				max = d
			}
		}

		if par.Part(1) {
			par.SubmitInt(1, d)
		}

		if par.Part(2) {
			par.SubmitInt(2, max)
		}
	}
}
