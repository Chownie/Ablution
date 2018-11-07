package ui

import "github.com/veandco/go-sdl2/sdl"

type Container interface {
	Add(Component)
	Remove(Component)
	SetDirty()
	Render() *sdl.Surface
}
