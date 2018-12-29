package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `91212129`, `9`)
	ins = ins.Test(2, `12131415`, `4`)

	for p := range ins.Advent(2017, 1) {
		fmt.Println(Brown("\n" + p.Name))
		in := p.Val

		sum1, sum2 := 0, 0
		for i := 0; i < len(in); i++ {
			if p.Part(1) && in[i] == in[(i+1)%len(in)] {
				sum1 += int(in[i] - '0')
			}
			if p.Part(2) && in[i] == in[(i+len(in)/2)%len(in)] {
				sum2 += int(in[i] - '0')
			}
		}

		if p.Part(1) {
			p.SubmitInt(1, sum1)
		}

		if p.Part(2) {
			p.SubmitInt(2, sum2)
		}
	}
}
