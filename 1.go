package main

import (
	"fmt"
	"metalim/advent/2017/lib/advent"

	. "github.com/logrusorgru/aurora"
)

func main() {
	for p := range advent.Inputs(2017, 1) {
		fmt.Println(Brown("\n" + p.Name))
		in := p.Val

		sum1, sum2 := 0, 0
		for i := 0; i < len(in); i++ {
			if in[i] == in[(i+1)%len(in)] {
				sum1 += int(in[i] - '0')
			}
			if in[i] == in[(i+len(in)/2)%len(in)] {
				sum2 += int(in[i] - '0')
			}
		}
		fmt.Println("part 1:", Green(sum1))
		fmt.Println("part 2:", Green(sum2))
	}
}
