package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Width           int32
	Height          int32
	Title           string
	internalSurface *sdl.Surface
	Components      []Component
}

func NewWindow(width int, height int, title string) *Window {
	surf, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	return &Window{int32(width), int32(height), title, surf, []Component{}}
}

func (w *Window) Add(comp Component) {
	w.Components = append(w.Components, comp)
	w.SetDirty()
}

func (w *Window) Remove(comp Component) {
	defer func() { w.SetDirty() }()
	for i, v := range w.Components {
		if v == comp {
			w.Components = append(w.Components[:i], w.Components[i+1:]...)
		}
	}
}

func (w *Window) SetDirty() {
	err := w.internalSurface.FillRect(&sdl.Rect{int32(0), int32(0),
		w.internalSurface.W, w.internalSurface.H}, 0xffffffff)
	if err != nil {
		panic(err)
	}
	for _, component := range w.Components {
		component.Render().Blit(component.Size(), w.internalSurface, component.Bounds())
	}
}

func (w *Window) Render() *sdl.Surface {
	return w.internalSurface
}
