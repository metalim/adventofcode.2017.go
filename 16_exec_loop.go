package main

import (
	"fmt"
	"metalim/advent/2017/lib/source"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func main() {
	var ins source.Inputs

	ins = ins.Test(3, `s1,x3/4,pe/b`, `baedc`, `ceadb`)

	for par := range ins.Advent(2017, 16) {
		fmt.Println(Brown("\n" + par.Name))
		spar := par.Split(",")

		sw := spar.Values
		ssn := spar.Ints()
		fmt.Println(len(sw), Black(sw[:3]).Bold(), Black(ssn[:3]).Bold())

		p := program{sw, ssn}
		init := "abcdefghijklmnop"
		part2Count := int(1e9)

		if strings.Contains(par.Name, "test") {
			init = init[:5]
			part2Count = 2
		}

		if par.Part(1) {
			par.Submit(1, p.exec(init))
		}

		if par.Part(2) {
			reg := init
			seen := map[string]int{}
			for i := 1; i <= part2Count; i++ {
				reg = p.exec(reg)
				if step, ok := seen[reg]; ok { // found a loop? -> warp to same instance before target.
					i = part2Count - (part2Count-step)%(i-step)
					seen = map[string]int{} // reset map
				}
				seen[reg] = i
			}
			par.Submit(2, string(reg))
		}
	}
}

type program struct {
	code []string
	args [][]int
}

func (p *program) exec(sreg string) string {
	mod := len(sreg)
	reg := []byte(sreg)
	for ip, c := range p.code {
		a := p.args[ip]
		switch c[0] {

		case 's':
			x := a[0] % mod
			reg = append(reg[mod-x:], reg[:mod-x]...)

		case 'x':
			reg[a[0]], reg[a[1]] = reg[a[1]], reg[a[0]]

		case 'p':
			s := string(reg)
			i := strings.IndexByte(s, c[1])
			j := strings.IndexByte(s, c[3])
			reg[i], reg[j] = reg[j], reg[i]

		}
	}
	return string(reg)
}
