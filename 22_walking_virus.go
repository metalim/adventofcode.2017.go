package main

import (
	"fmt"
	"log"
	"metalim/advent/2017/lib/field"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `..#
#..
...`

func main() {
	var ins source.Inputs

	ins = ins.Test(3, test1, `5587`, `2511944`)

	for par := range ins.Advent(2017, 22) {
		fmt.Println(Brown("\n" + par.Name))
		smap := par.Val

		if par.Part(1) {
			f := &field.Slice{}
			f.SetDefault('.')
			field.FillFromString(f, field.Pos{}, smap)

			b := f.Bounds()
			start := field.Pos{b.Dx() / 2, b.Dy() / 2}

			var infected, istep int

			canStepOn := func(field.Pos) bool { return true }
			stepOn := func(p field.Pos) bool {
				istep++
				return true
			}
			getDirs := func(p field.Pos, d field.Dir8) int {
				v0 := f.Get(p)
				v := int('.')
				if v0 == '.' {
					v = '#'
					infected++
				}
				f.Set(p, v)

				if istep == 1e4 {
					return 0
				}
				if v0 == '#' {
					return 1 << ((d + 2) & 7) // prev infected -> right
				}
				return 1 << ((d + 6) & 7) // prev clean -> left
			}
			field.Walk(start, field.Dir80N, canStepOn, stepOn, getDirs)

			par.SubmitInt(1, infected)
		}

		if par.Part(2) {
			f := &field.Slice{}
			f.SetDefault('.')
			field.FillFromString(f, field.Pos{}, smap)

			b := f.Bounds()
			start := field.Pos{b.Dx() / 2, b.Dy() / 2}

			var infected, istep int

			canStepOn := func(field.Pos) bool { return true }
			stepOn := func(p field.Pos) bool {
				istep++
				return true
			}
			getDirs := func(p field.Pos, d field.Dir8) int {
				v0 := f.Get(p)
				var v, dirs int
				switch v0 {
				case '.':
					v = 'W'
					dirs = 1 << ((d + 6) & 7)
				case 'W':
					v = '#'
					infected++
					dirs = 1 << d
				case '#':
					v = 'F'
					dirs = 1 << ((d + 2) & 7)
				case 'F':
					v = '.'
					dirs = 1 << ((d + 4) & 7)
				default:
					log.Fatal("unkown cell value", string(v0))
				}
				f.Set(p, v)

				if istep == 1e7 {
					return 0
				}
				return dirs
			}
			field.Walk(start, field.Dir80N, canStepOn, stepOn, getDirs)

			par.SubmitInt(2, infected)
		}
	}
}
