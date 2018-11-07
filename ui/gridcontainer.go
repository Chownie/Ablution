package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GridContainer struct {
	internalSurface *sdl.Surface
	*EventManager
	*BaseContainer
}

func NewGridContainer(x int, y int, width int, height int) *GridContainer {
	surf, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	return &GridContainer{surf, NewEventManager(), NewBaseContainer(x, y, width, height, false)}
}

func (g *GridContainer) SetDirty() {
	err := g.internalSurface.FillRect(&sdl.Rect{int32(0), int32(0),
		g.internalSurface.W, g.internalSurface.H}, 0xffffffff)
	if err != nil {
		panic(err)
	}
	for _, component := range g.Components {
		component.Render().Blit(component.Size(), g.internalSurface, component.Bounds())
	}
}

func (g *GridContainer) Render() *sdl.Surface {
	return g.internalSurface
}
