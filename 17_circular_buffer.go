package main

import (
	"fmt"
	"metalim/advent/2017/lib/debug"
	"metalim/advent/2017/lib/source"
	"time"

	. "github.com/logrusorgru/aurora"
)

func main() {
	// source.Dry()

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
			n := &node{val: 0}
			n.next = n
			var i int
			var stop bool
			go func() {
				for !stop {
					debug.Log(Black(i).Bold())
					time.Sleep(time.Second)
				}
			}()
			var after0 int
			for i = 1; i <= 5e7; i++ {
				n = n.skip(step)
				if n.val == 0 {
					after0 = i
				}
				prev := n
				n = &node{val: i, next: prev.next}
				prev.next = n
			}
			stop = true
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
