package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `3`, `638`)

	for p := range ins.Advent(2017, 17) {
		fmt.Println(Brown("\n" + p.Name))
		step := p.Ints()[0]
		fmt.Println(Black(step).Bold())

		if p.Part(1) {
			n := &node{val: 0}
			n.next = n
			for i := 1; i <= 2017; i++ {
				n = n.skip(step)
				prev := n
				n = &node{val: i, next: prev.next}
				prev.next = n
			}
			p.SubmitInt(1, n.next.val)
		}

		if p.Part(2) {
			var after0, pos0, pos int
			length := 1
			for i := 1; i <= 5e7; i++ {
				pos = (pos + step) % length
				if pos < pos0 {
					pos0++
				} else if pos == pos0 {
					after0 = i
				}
				pos++
				length++
			}
			p.SubmitInt(2, after0)
		}
	}
}

type node struct {
	val  int
	next *node
}

func (n *node) skip(i int) *node {
	for ; i > 0; i-- {
		n = n.next
	}
	return n
}
