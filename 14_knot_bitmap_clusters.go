package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"metalim/advent/2017/lib/union"

	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `flqrgnkx`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `8108`)
	ins = ins.Test(2, test1, `1242`)

	for p := range ins.Advent(2017, 14) {
		fmt.Println(Brown("\n" + p.Name))
		s := p.Val
		fmt.Println(Black(s).Bold())

		const dim = 128

		if p.Part(1) {
			var set int
			for i := 0; i < dim; i++ {
				bytes := knotHash(s + "-" + strconv.Itoa(i))
				for _, b := range bytes {
					for b > 0 {
						if b&1 != 0 {
							set++
						}
						b >>= 1
					}
				}
			}
			p.SubmitInt(1, set)
		}

		if p.Part(2) {
			var prev [dim]bool // we don't need whole bitmap, just the previous row, and union set.
			u := union.New()
			for y := 0; y < dim; y++ {
				bytes := knotHash(s + "-" + strconv.Itoa(y))
				var row [dim]bool
				for i, b := range bytes {
					for j := 0; j <= 7; j++ {
						if (b>>uint(7-j))&1 != 0 {
							x := i*8 + j
							k := dim*y + x
							row[x] = true
							u.Link(k) // standalone bit also counts as a region.
							if x > 0 && row[x-1] {
								u.Link(k, k-1)
							}
							if y > 0 && prev[x] {
								u.Link(k, k-dim)
							}
						}
					}
				}
				prev = row
			}
			p.SubmitInt(2, len(u.Unions))
		}

	}
}

func knotHash(s string) (out [16]byte) {
	sn := []byte(s)
	sn = append(sn, 17, 31, 73, 47, 23)

	const dim = 256
	var list [dim]byte
	for i := range list {
		list[i] = byte(i)
	}

	var pos, skip int
	for round := 0; round < 64; round++ {
		for _, n := range sn {
			for i := 0; i < int(n/2); i++ { // reverse chunk of the circle.
				a := (pos + i) % dim
				b := (pos + int(n) - 1 - i) % dim
				list[a], list[b] = list[b], list[a]
			}
			pos = (pos + int(n) + skip) % dim
			skip++
		}
	}

	// sparse -> dense
	for i := 0; i < 16; i++ {
		for _, j := range list[i*16 : i*16+16] {
			out[i] ^= j
		}
	}
	return
}
