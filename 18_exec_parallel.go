package main

import (
	"fmt"
	"metalim/advent/2017/lib/debug"
	"metalim/advent/2017/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

var test2 = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `4`)
	ins = ins.Test(2, test2, `3`)

	for par := range ins.Advent(2017, 18) {
		fmt.Println(Brown("\n" + par.Name))
		parl := par.Lines()
		ssw := parl.Words()
		ssn := parl.Ints() // unreliable to pair with Words, as it seems.
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		if par.Part(1) {
			reg := map[string]int{}
			var ip, snd, rcv int
		LOOP1:
			for ip >= 0 && ip < len(ssw) {
				op := ssw[ip]
				arg := ssn[ip]
				val := func(i int) int {
					if len(arg) > 0 {
						return arg[0]
					}
					return reg[op[i]]
				}
				switch op[0] {
				case "set":
					reg[op[1]] = val(2)
				case "add":
					reg[op[1]] += val(2)
				case "mul":
					reg[op[1]] *= val(2)
				case "mod":
					reg[op[1]] %= val(2)
				case "jgz":
					if reg[op[1]] > 0 { // bug here. But works for part 1.
						ip += val(2) - 1
					}
				case "snd":
					snd = val(1)
				case "rcv":
					if snd != 0 {
						rcv = snd
						break LOOP1
					}
				}
				ip++
			}

			par.SubmitInt(1, rcv)
		}

		if par.Part(2) {
			var sent1 int
			qs := [2][]int{}
			ips := [2]int{}
			regs := [2]map[string]int{{"p": 0}, {"p": 1}}
			waiting := [2]bool{}

		LOOP2:
			for {

				if waiting[0] && len(qs[0]) == 0 && waiting[1] && len(qs[1]) == 0 {
					debug.Log("deadlock")
					break LOOP2
				}

			NEXTP:
				for p := 0; p <= 1; p++ {
					ip := ips[p]
					reg := regs[p]
					op := ssw[ip]

					val := func(i int) int {
						if v, err := strconv.Atoi(op[i]); err == nil { // can't rely on .Ints(), as there can be 2 numbers.
							return v
						}
						return reg[op[i]]
					}

					switch op[0] {
					case "set":
						reg[op[1]] = val(2)
					case "add":
						reg[op[1]] += val(2)
					case "mul":
						reg[op[1]] *= val(2)
					case "mod":
						reg[op[1]] %= val(2)
					case "jgz":
						if val(1) > 0 { // first argument can be number, dammit !!!
							ips[p] += val(2)
							continue NEXTP
						}

					case "snd":
						if p == 1 {
							sent1++
						}
						qs[1-p] = append(qs[1-p], val(1))

					case "rcv":
						if len(qs[p]) == 0 {
							waiting[p] = true
							continue NEXTP
						}
						waiting[p] = false
						reg[op[1]] = qs[p][0]
						qs[p] = qs[p][1:]

					default:
						panic("unknown instruction: " + fmt.Sprint(op))
					}

					ips[p] = ip + 1
				}
			}

			par.SubmitInt(2, sent1)
		}
	}
}

func exec() {

}
