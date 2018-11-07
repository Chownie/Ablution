package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GridContainer struct {
	internalSurface *sdl.Surface
	Components      []Component
	*EventManager
	*Transform
}

func NewGridContainer(x int, y int, width int, height int) *GridContainer {
	surf, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	return &GridContainer{surf, []Component{}, NewEventManager(), NewTransform(x, y, width, height, true)}
}

func (g *GridContainer) Add(Component) {

}

func (g *GridContainer) Remove(Component) {

}

func (g *GridContainer) SetDirty() {

}

func (g *GridContainer) Render() *sdl.Surface {
	return g.internalSurface
}
