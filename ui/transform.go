package ui

import "github.com/veandco/go-sdl2/sdl"

type Transform struct {
	autosize bool
	minimumW int
	minimumH int
	maximumW int
	maximumH int
	*sdl.Rect
}

func NewTransform(x int, y int, w int, h int, autosize bool) *Transform {
	return &Transform{autosize, 0, 0, 0, 0, &sdl.Rect{
		int32(x),
		int32(y),
		int32(w),
		int32(h),
	}}
}

func (t *Transform) SetRange(minw int, maxw int, minh int, maxh int) {
	t.minimumW = minw
	t.minimumH = minh
	t.maximumW = maxw
	t.maximumH = maxh
}

func (t *Transform) Range(min int, input int, max int) int32 {
	if input < min {
		return int32(min)
	}
	if input > max {
		return int32(max)
	}
	return int32(input)
}

func (t *Transform) Resize(w int, h int) {
	if !t.autosize {
		t.W = int32(w)
		t.H = int32(h)
	}
	t.W = t.Range(t.minimumW, w, t.maximumW)
	t.H = t.Range(t.minimumH, h, t.maximumH)
}

func (t *Transform) Contains(x int32, y int32) bool {
	pt := sdl.Point{x, y}
	if pt.InRect(t.Bounds()) {
		return true
	}
	return false
}

func (t *Transform) Size() *sdl.Rect {
	return &sdl.Rect{int32(0), int32(0), t.W, t.H}
}

func (t *Transform) Bounds() *sdl.Rect {
	return &sdl.Rect{t.X, t.Y, t.W, t.H}
}
