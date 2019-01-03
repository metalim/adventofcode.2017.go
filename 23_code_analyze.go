package main

import (
	"fmt"
	"log"
	"math"
	"metalim/advent/2017/lib/source"
	"strconv"

	. "github.com/logrusorgru/aurora"
)

var test1 = ``

func main() {
	var ins source.Inputs

	for par := range ins.Advent(2017, 23) {
		fmt.Println(Brown("\n" + par.Name))
		ssw := par.Lines().Words()
		fmt.Println(len(ssw), Black(ssw[0]).Bold())

		if par.Part(1) {
			reg := map[string]int{}
			var ip, muls int

		NEXT:
			for ip >= 0 && ip < len(ssw) {
				op := ssw[ip]
				val := func(i int) int {
					if v, err := strconv.Atoi(op[i]); err == nil {
						return v
					}
					return reg[op[i]]
				}
				switch op[0] {
				case "set":
					reg[op[1]] = val(2)
				case "sub":
					reg[op[1]] -= val(2)
				case "mul":
					muls++
					reg[op[1]] *= val(2)
				case "jnz":
					if val(1) != 0 {
						ip += val(2)
						continue NEXT
					}
				default:
					log.Fatal("Unknown instruction", op)
				}
				ip++
			}
			par.SubmitInt(1, muls)
		}

		if par.Part(2) {

			b, _ := strconv.Atoi(ssw[0][2])
			par.SubmitInt(2, step6(1, b))
		}
	}
}

func step6(a, b int) (h int) {
	var c, d, e int

	// b = seed, min.
	c = b // c = max.
	if a != 0 {
		b = b*100 + 100000
		c = b + 17000
	}

	for ; b <= c; b += 17 {
		// b is prime? -> h++
		sq := int(math.Sqrt(float64(b)))
		for d = 2; d <= sq; d++ {
			e = b / d
			if d*e == b {
				h++
				break
			}
		}
	}
	return h
}
