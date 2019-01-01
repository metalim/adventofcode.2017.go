package field

import "strings"

type map1d map[int]Cell
type map2d map[int]map1d

// Map z
type Map struct {
	field2d
	m map2d
}

// Get cell.
func (f *Map) Get(p Pos) Cell {
	if r, ok := f.m[p.Y]; ok {
		return r[p.X]
	}
	return f.def
}

// Set cell.
func (f *Map) Set(p Pos, c Cell) {
	if f.m == nil {
		f.m = map2d{}
	}
	if _, ok := f.m[p.Y]; !ok {
		f.m[p.Y] = map1d{}
	}
	if !p.In(f.b) {
		f.b = f.b.Union(Rect{p, p.Add(Pos{1, 1})})
	}
	f.m[p.Y][p.X] = c
}

// FillFromString from start position.
func (f *Map) FillFromString(start Pos, s string) {
	for y, l := range strings.Split(s, "\n") {
		for x, r := range l {
			f.Set(start.Add(Pos{x, y}), Cell(r))
		}
	}
}
