package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

type node struct {
	name                string
	weight, weightTotal int
	unbalanced          bool
	kidNames            []string
	kids                []*node
}

func calc(n *node) int {
	n.weightTotal = n.weight
	if len(n.kids) > 0 {
		for _, k := range n.kids {
			n.weightTotal += calc(k)
		}
		w := n.kids[0].weightTotal
		for _, k := range n.kids[1:] {
			if w != k.weightTotal {
				n.unbalanced = true
				break
			}
		}
	}
	return n.weightTotal
}

func main() {
	for p := range source.Test(1, test1, `tknk`).Test(2, test1, `60`).Advent(2017, 7) {
		fmt.Println(Brown("\n" + p.Name))
		ssw := p.Lines().WordsTrim("()->, ")
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		// parse flat.
		nodes := map[string]*node{}
		kidParents := map[string]*node{} // map kid names to parents.
		for _, l := range ssw {
			w, _ := strconv.Atoi(l[1])
			kids := []string{}
			if len(l) > 3 {
				kids = l[3:]
			}
			n := &node{name: l[0], weight: w, kidNames: kids}
			nodes[n.name] = n
			for _, c := range kids {
				kidParents[c] = n
			}
		}

		// grow tree.
		for k, p := range kidParents {
			p.kids = append(p.kids, nodes[k])
			delete(nodes, k)
		}

		// select root.
		var root *node
		for _, n := range nodes {
			root = n
		}

		if p.Part(1) {
			p.Submit(1, root.name)
		}

		if p.Part(2) {
			// calculate weights and balances
			calc(root)

			// find last unbalanced node
			un := root
		LOOP:
			for {
				for _, k := range un.kids {
					if k.unbalanced {
						un = k
						continue LOOP
					}
				}
				// all kids are balanced, but node is not, so it's one of kids.
				break
			}

			// kids weight counts
			ws := map[int]int{}
			for _, k := range un.kids {
				ws[k.weightTotal]++
			}

			// incorrect and correct weight
			var w1, w2 int
			for w, count := range ws {
				if count == 1 {
					w1 = w
				} else {
					w2 = w
				}
			}

			// and finally, select incorrect
			for _, n := range un.kids {
				if n.weightTotal == w1 {
					p.SubmitInt(2, n.weight-w1+w2)
					break
				}
			}
		}
	}
}
