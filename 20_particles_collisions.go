package main

import (
	"fmt"
	"metalim/advent/2017/lib/numbers"
	"metalim/advent/2017/lib/source"

	. "github.com/logrusorgru/aurora"
)

var test1 = `p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>`

var test2 = `p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>`

func main() {
	var ins source.Inputs

	ins = ins.Test(1, test1, `0`)
	ins = ins.Test(2, test2, `1`)

	for par := range ins.Advent(2017, 20) {
		fmt.Println(Brown("\n" + par.Name))
		ps := par.Lines().Ints()
		fmt.Println(len(ps), Black(ps[:2]).Bold())

		if par.Part(1) {
			for _, p := range ps {
				for j := 0; j < 500; j++ {
					p[3] += p[6]
					p[4] += p[7]
					p[5] += p[8]

					p[0] += p[3]
					p[1] += p[4]
					p[2] += p[5]
				}
			}
			mind := 999999999
			mini := -1
			for i, p := range ps {
				d := manh(p)
				if mind > d {
					mind = d
					mini = i
				}

			}
			par.SubmitInt(1, mini)
		}

		if par.Part(2) {
			ps := par.Lines().Ints() // reset input

			for step := 0; step < 500; step++ {

				// collide all
				removed := map[int]bool{}
				for i, p1 := range ps {
					for j, p2 := range ps[i+1:] {
						if eq(p1[:3], p2[:3]) {
							removed[i] = true
							removed[j+i+1] = true
						}
					}
				}

				// remove collided
				var out int
				for i, p := range ps {
					if !removed[i] {
						ps[out] = p
						out++
					}
				}
				ps = ps[:out]

				// move all
				for _, p := range ps {
					p[3] += p[6]
					p[4] += p[7]
					p[5] += p[8]

					p[0] += p[3]
					p[1] += p[4]
					p[2] += p[5]
				}
			}
			par.SubmitInt(2, len(ps))
		}

	}
}

func eq(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func manh(p []int) int {
	return numbers.Abs(p[0]) + numbers.Abs(p[1]) + numbers.Abs(p[2])
}
