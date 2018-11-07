package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	parent          *Window
	internalSurface *sdl.Surface
	backColor       *sdl.Color
	*EventManager
	*Transform
}

/*func NewImage(x int, y int) *Image {
	surf, err := sdl.CreateRGBSurface(0, tex.W, tex.H, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	image := Image{Drawing, nil, font, surf, tex, NewEventManager(), NewTransform(x, y, int(tex.W), int(tex.H), true)}
	return &image
}
*/
