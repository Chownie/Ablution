package ui

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Container interface {
	Add(Component)
	Remove(Component)
	SetDirty()
	Render() *sdl.Surface
}

type BaseContainer struct {
	internalSurface *sdl.Surface
	Components      []Component
	parent          Container
	*EventManager
	*Transform
}

func NewBaseContainer(x int, y int, width int, height int, autosize bool) *BaseContainer {
	surf, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	return &BaseContainer{surf, []Component{}, nil, NewEventManager(), NewTransform(x, y, width, height, autosize)}
}

func (b *BaseContainer) Add(component Component) {
	b.Components = append(b.Components, component)
}

func (b *BaseContainer) Remove(component Component) {
	fmt.Print("Adding:")
	for i, v := range b.Components {
		if v == component {
			b.Components = append(b.Components[:i], b.Components[i+1:]...)
			fmt.Println(len(b.Components))
		}
	}
}

func (b *BaseContainer) SetDirty() {
	err := b.internalSurface.FillRect(&sdl.Rect{int32(0), int32(0),
		b.internalSurface.W, b.internalSurface.H}, 0xffffffff)
	if err != nil {
		panic(err)
	}
	for _, component := range b.Components {
		component.Render().Blit(component.Size(), b.internalSurface, component.Bounds())
	}
}

func (b *BaseContainer) SetParent(container Container) {
	b.parent = container
	b.SetDirty()
}

func (g *BaseContainer) Render() *sdl.Surface {
	return g.internalSurface
}
