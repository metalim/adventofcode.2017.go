package main

import (
	"fmt"
	"log"
	"metalim/advent/2017/lib/field"
	"metalim/advent/2017/lib/source"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var test1 = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `12`)

	for par := range ins.Advent(2017, 21) {
		fmt.Println(Brown("\n" + par.Name))
		sss := par.Lines().Split(" => ")
		fmt.Println("rules:", Black(len(sss)).Bold())

		rules := map[string]string{}
		for _, ss := range sss {
			skey := strings.Replace(ss[0], "/", "\n", -1)
			sval := strings.Replace(ss[1], "/", "\n", -1)
			rules[skey] = sval

			// rotate x3, flip, rotate x3
			key := &field.Slice{}
			field.FillFromString(key, field.Pos{}, skey)
			rect := key.Bounds()

			for f := 0; f < 2; f++ {
				for r := 0; r < 3; r++ {
					field.Rotate90(key, rect)
					rules[field.ToString(key, rect)] = sval
				}
				field.FlipVert(key, rect) // there was bug here: flipped unrotated key was not added.
				rules[field.ToString(key, rect)] = sval
			}
		}

		if par.Part(1) {
			steps := 5
			if strings.Contains(par.Name, "test") {
				steps = 2
			}
			steps = steps
			par.SubmitInt(1, iterate(rules, steps))
		}

		if par.Part(2) {
			par.SubmitInt(2, iterate(rules, 18))
		}
	}
}

func iterate(rules map[string]string, steps int) (out int) {
	cur := &field.Slice{}
	field.FillFromString(cur, field.Pos{}, ".#.\n..#\n###")

	for step := 1; step <= steps; step++ {
		b := cur.Bounds()
		var d int
		switch {
		case b.Dx()%2 == 0:
			d = 2
		case b.Dx()%3 == 0:
			d = 3
		default:
			panic("invalid dimensions")
		}
		next := &field.Slice{}

		for y := b.Min.Y; y < b.Max.Y; y += d {
			for x := b.Min.X; x < b.Max.X; x += d {
				dx := x / d * (d + 1)
				dy := y / d * (d + 1)
				key := field.ToString(cur, field.Rect{field.Pos{x, y}, field.Pos{x + d, y + d}})
				val, ok := rules[key]
				if !ok {
					log.Fatal("step ", step, " can't find rule for:\n", key)
				}
				field.FillFromString(next, field.Pos{dx, dy}, val)
			}
		}
		cur = next
	}

	// count cells.
	b := cur.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if cur.Get(field.Pos{x, y}) == '#' {
				out++
			}
		}
	}
	return out
}
