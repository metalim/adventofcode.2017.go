package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `65 asdf blabla 8921`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `588`)
	ins = ins.Test(2, test1, `309`)

	for p := range ins.Advent(2017, 15) {
		fmt.Println(Brown("\n" + p.Name))
		sn := p.Ints()
		fmt.Println(Black(sn).Bold())

		const fa = 16807
		const fb = 48271
		const mod = 2147483647
		if p.Part(1) {
			var pairs int
			a := sn[0]
			b := sn[1]
			for i := 0; i < 4e7; i++ {
				a = a * fa % mod
				b = b * fb % mod
				if a&0xffff == b&0xffff {
					pairs++
				}
			}
			p.SubmitInt(1, pairs)
		}

		if p.Part(2) {
			var pairs int
			a := sn[0]
			b := sn[1]
			for i := 0; i < 5e6; i++ {
				for ok := false; !ok; ok = a%4 == 0 {
					a = a * fa % mod
				}
				for ok := false; !ok; ok = b%8 == 0 {
					b = b * fb % mod
				}
				if a&0xffff == b&0xffff {
					pairs++
				}
			}
			p.SubmitInt(2, pairs)
		}
	}
}
