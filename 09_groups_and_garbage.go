package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `{{{<{}>},{},{{}}}}`

var test2 = `<{o"i!a,<{i<a>`

func main() {
	for p := range source.Test(1, test1, `16`).Test(2, test2, `10`).Advent(2017, 9) {
		fmt.Println(Brown("\n" + p.Name))
		s := p.Val
		fmt.Println(len(s), Black(s[:14]).Bold())

		var score, level, garbage int
		var inGarbage bool
		for i := 0; i < len(s); i++ {
			if inGarbage {
				switch s[i] {
				case '>':
					inGarbage = false
				case '!':
					i++
				default:
					garbage++
				}
				continue
			}
			switch s[i] {
			case '{':
				level++
				score += level
			case '}':
				level--
			case '<':
				inGarbage = true
			}
		}

		if p.Part(1) {
			p.SubmitInt(1, score)
		}

		if p.Part(2) {
			p.SubmitInt(2, garbage)
		}
	}
}
