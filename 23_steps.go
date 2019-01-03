import (
	"math"
	"metalim/advent/2017/lib/debug"
)

////////////////////////////////////////////////////////////////////////

func step0(a int) (h int) {
	var b, c, d, e, f, g int

	// 0: set b 57
	b = 57
	// 1: set c b
	c = b
	// 2: jnz a 2
	if a != 0 {
		goto L4
	}
	// 3: jnz 1 5
	goto L8
L4:
	// 4: mul b 100
	b *= 100
	// 5: sub b -100000
	b += 100000
	// 6: set c b
	c = b
	// 7: sub c -17000
	c += 17000
L8:
	// 8: set f 1
	f = 1
	// 9: set d 2
	d = 2
L10:
	// 10: set e 2
	e = 2
L11:
	// 11: set g d
	g = d
	// 12: mul g e
	g *= e
	// 13: sub g b
	g -= b
	// 14: jnz g 2
	if g != 0 {
		goto L16
	}
	// 15: set f 0
	f = 0
L16:
	// 16: sub e -1
	e++
	// 17: set g e
	g = e
	// 18: sub g b
	g -= b
	// 19: jnz g -8
	if g != 0 {
		goto L11
	}
	// 20: sub d -1
	d++
	// 21: set g d
	g = d
	// 22: sub g b
	g -= b
	// 23: jnz g -13
	if g != 0 {
		goto L10
	}
	// 24: jnz f 2
	if f != 0 {
		goto L26
	}
	// 25: sub h -1
	h++
L26:
	// 26: set g b
	g = b
	// 27: sub g c
	g -= c
	// 28: jnz g 2
	if g != 0 {
		goto L30
	}
	// 29: jnz 1 3
	goto L32
L30:
	// 30: sub b -17
	b += 17
	// 31: jnz 1 -23
	goto L8

L32:
	return h
}

////////////////////////////////////////////////////////////////////////

func step1(a, b int) (h int) {
	var c, d, e, f, g int
	var muls int

	c = b
	if a != 0 {
		goto L4
	}
	goto L8
L4:
	b *= 100
	muls++
	b += 100000
	c = b
	c += 17000
L8:
	f = 1
	d = 2
L10:
	e = 2
L11:
	g = d
	g *= e
	muls++
	g -= b
	if g != 0 {
		goto L16
	}
	f = 0
L16:
	e++
	g = e
	g -= b
	if g != 0 {
		goto L11
	}
	d++
	g = d
	g -= b
	if g != 0 {
		goto L10
	}
	if f != 0 {
		goto L26
	}
	h++
L26:
	g = b
	g -= c
	if g != 0 {
		goto L30
	}
	goto L32
L30:
	b += 17
	goto L8

L32:
	debug.Log(muls)
	return h
}

////////////////////////////////////////////////////////////////////////

func step2(a, b int) (h int) {
	var c, d, e, f, g int
	var muls int

	c = b
	if a == 0 {
		goto L8
	}
	b *= 100
	muls++
	b += 100000
	c = b
	c += 17000
L8:
	f = 1
	d = 2
L10:
	e = 2
L11:
	g = d
	g *= e
	muls++
	g -= b
	if g != 0 {
		goto L16
	}
	f = 0
L16:
	e++
	g = e
	g -= b
	if g != 0 {
		goto L11
	}
	d++
	g = d
	g -= b
	if g != 0 {
		goto L10
	}
	if f != 0 {
		goto L26
	}
	h++
L26:
	g = b
	g -= c
	if g == 0 {
		goto L32
	}
	b += 17
	goto L8

L32:
	debug.Log(muls)
	return h
}

////////////////////////////////////////////////////////////////////////

func step3(a, b int) (h int) {
	var c, d, e, f, g int
	var muls int

	c = b
	if a != 0 {
		muls++
		b = b*100 + 100000
		c = b + 17000
	}
L8:
	f = 1
	d = 2
L10:
	e = 2
L11:
	muls++
	g = d*e - b
	if g == 0 {
		f = 0
	}
	e++
	g = e - b
	if g != 0 {
		goto L11
	}
	d++
	g = d - b
	if g != 0 {
		goto L10
	}
	if f == 0 {
		h++
	}
	g = b - c
	if g == 0 {
		goto L32
	}
	b += 17
	goto L8

L32:
	debug.Log(muls)
	return h
}

////////////////////////////////////////////////////////////////////////

func step4(a, b int) (h int) {
	var c, d, e, f, g int
	var muls int

	// b = seed
	c = b
	if a != 0 {
		muls++
		b = b*100 + 100000
		c = b + 17000
	}
	debug.Log(a, b, c)

	for {
		f = 1
		d = 2
		for g0 := false; !g0; g0 = d == b {
			e = 2
			for g00 := false; !g00; g00 = e == b {
				muls++
				if d*e == b {
					f = 0
				}
				e++
			}
			d++
		}
		if f == 0 {
			h++
		}
		g = b - c
		if g == 0 {
			break
		}
		b += 17
	}
	debug.Log(muls)
	return h
}

////////////////////////////////////////////////////////////////////////

func step5(a, b int) (h int) {
	var c, d, e, f int
	var muls int

	// b = seed
	c = b
	if a != 0 {
		muls++
		b = b*100 + 100000
		c = b + 17000
	}
	debug.Log(a, b, c)

	for ; ; b += 17 {
		f = 1
		// LOOP2:
		for d = 2; d != b; d++ {
			for e = 2; e != b; e++ {
				muls++
				if d*e == b {
					f = 0
					// break LOOP2
				}
			}
		}
		if f == 0 {
			h++
		}
		if b == c {
			break
		}
	}
	debug.Log(muls)
	return h
}

////////////////////////////////////////////////////////////////////////

func step5(a, b int) (h int) {
	var c, d, e, f int

	// b = seed
	c = b
	if a != 0 {
		b = b*100 + 100000
		c = b + 17000
	}

	for ; b <= c; b += 17 {
		f = 1
		// f := is b prime?
		sq := int(math.Sqrt(float64(b)))
		for d = 2; d <= sq; d++ {
			e = b / d
			if d*e == b {
				f = 0
				break
			}
		}
		if f == 0 {
			h++
		}
	}
	return h
}

////////////////////////////////////////////////////////////////////////

func step6(a, b int) (h int) {
	var c, d, e int

	// b = seed, min.
	c = b // c = max.
	if a != 0 {
		b = b*100 + 100000
		c = b + 17000
	}

	for ; b <= c; b += 17 {
		// if b is prime? then h++
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

////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////
