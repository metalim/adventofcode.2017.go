package main

import (
	"fmt"
	"metalim/advent/2017/lib/graph"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `31`)

	for par := range ins.Advent(2017, 24) {
		fmt.Println(Brown("\n" + par.Name))
		ssn := par.Lines().Ints()
		fmt.Println(len(ssn), Black(ssn[0]).Bold())

		g := graph.NewGraph()
		for _, l := range ssn {
			g.Link(g.Node(l[0]), g.Node(l[1])).Data = l[0] + l[1]
		}

		if par.Part(1) {
			var maxSum int
			g.IteratePathsFrom(g.Node(0), func(path graph.Path) bool {
				var sum int
				for _, ln := range path {
					sum += ln.Data.(int)
				}
				if maxSum < sum {
					maxSum = sum
				}
				return true
			})
			par.SubmitInt(1, maxSum)
		}

		if par.Part(2) {
			var maxLen, maxLenSum int
			g.IteratePathsFrom(g.Node(0), func(path graph.Path) bool {
				if maxLen > len(path) {
					return true
				}
				maxLen = len(path)
				var sum int
				for _, ln := range path {
					sum += ln.Data.(int)
				}
				if maxLenSum < sum {
					maxLenSum = sum
				}
				return true
			})
			par.SubmitInt(2, maxLenSum)
		}

	}
}

////////////////////////////////////////////////////////////////////////
