package main

import (
	"fmt"
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

		g := NewGraph()
		for _, l := range ssn {
			g.Link(l[0]+l[1], g.Node(l[0]), g.Node(l[1]))
		}

		if par.Part(1) {
			var maxSum int
			g.IteratePathsFrom(g.Node(0), func(path []*Link) bool {
				var sum int
				for _, ln := range path {
					sum += ln.weight
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
			g.IteratePathsFrom(g.Node(0), func(path []*Link) bool {
				if maxLen > len(path) {
					return true
				}
				maxLen = len(path)
				var sum int
				for _, ln := range path {
					sum += ln.weight
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

type idt int

// Graph .
type Graph struct {
	nodes    map[idt]*Node
	links    map[idt]*Link
	idn, idl idt
}

// NewGraph creates maps.
func NewGraph() *Graph {
	return &Graph{nodes: map[idt]*Node{}, links: map[idt]*Link{}}
}

// Link - not using name "edge", as the name is stupid.
type Link struct {
	id     idt
	weight int
	// nodes  [2]*Node // we need it for g.Unlink, and if we're interested in nodes in general.
	// or
	// nodes map[idt]*nodes // can be done for multilinks ("hyperedges" in "hypergraphs").
}

type linksMap map[*Link]*Node

// Node - not using name "vertice", for same reasons.
type Node struct {
	id    idt
	links linksMap
}

// Node returns existing node or creates new one.
func (g *Graph) Node(id int) *Node {
	n, ok := g.nodes[idt(id)]
	if !ok {
		n = &Node{id: g.idn, links: linksMap{}}
		g.idn++
		g.nodes[idt(id)] = n
	}
	return n
}

// Link bidirectional.
func (g *Graph) Link(weight int, a, b *Node) *Link {
	ln := &Link{id: g.idl, weight: weight}
	g.links[g.idl] = ln
	g.idl++
	// every to every
	a.links[ln] = b // works for self-linking loops,
	b.links[ln] = a // as this just replaces with same value.
	return ln
}

func (n *Node) copyLinks() linksMap {
	out := linksMap{}
	for l, n := range n.links {
		out[l] = n
	}
	return out
}

// IteratePathsFrom starts from specified node and iterates all possible paths in DFS.
// Path is followed only if callback returns true.
func (g *Graph) IteratePathsFrom(start *Node, fn func([]*Link) bool) {
	lns := start.copyLinks()
	used := map[*Link]bool{}
	var path []*Link
	var queues []linksMap // expensive.

APPEND:
	for {
		for ln, n2 := range lns {
			delete(lns, ln) // it's a queue in map type :-)
			if used[ln] {   // skip already used links
				continue
			}
			next := append(path, ln)
			if !fn(next) { // follow this path only if callback returns true.
				continue
			}
			path = next
			used[ln] = true
			queues = append(queues, lns)
			lns = n2.copyLinks()
			continue APPEND // with new set of links
		}
		// no more appends -> pop last link, and continue with remaining links.
		if len(queues) == 0 { // nothing to pop -> exit.
			break
		}
		lns, queues = queues[len(queues)-1], queues[:len(queues)-1]
		var released *Link
		released, path = path[len(path)-1], path[:len(path)-1]
		delete(used, released)
	}
}
