package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"sort"

	. "github.com/logrusorgru/aurora"
)

var test1 = `aa bb cc dd aa`
var test2 = `aa bb cc dd aaa`
var test3 = `abcde xyz ecdab`

func main() {
	for p := range source.Test(1, test1, `0`).Test(1, test2, `1`).Test(2, test3, `0`).Advent(2017, 4) {
		fmt.Println(Brown("\n" + p.Name))
		ssw := p.Lines().Words()
		fmt.Println(len(ssw), "lines", Black(ssw[0]).Bold())
		if p.Part(1) {
			var count int
		LINES:
			for _, l := range ssw {
				words := map[string]bool{}
				for _, w := range l {
					if words[w] {
						continue LINES
					}
					words[w] = true
				}
				count++
			}
			p.SubmitInt(1, count)
		}
		if p.Part(2) {
			var count int
		LINES2:
			for _, l := range ssw {
				words := map[string]bool{}
				for _, w := range l {
					rs := []rune(w)
					sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
					w = string(rs)
					if words[w] {
						continue LINES2
					}
					words[w] = true
				}
				count++
			}
			p.SubmitInt(2, count)
		}
	}
}
