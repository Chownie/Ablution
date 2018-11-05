package ui

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Component interface {
	Render() *sdl.Surface
	Size() *sdl.Rect
	Bounds() *sdl.Rect
	Contains(int32, int32) bool
	SetParent(*Window)
	SetDirty()
	Register(Event, func(EventData))
	Trigger(Event, EventData) bool
}

func positionMiddle(object *sdl.Surface, boundary *sdl.Surface) *sdl.Rect {
	w := object.W
	h := object.H
	x := boundary.W/2 - w/2
	y := boundary.H/2 - h/2
	return &sdl.Rect{x, y, w, h}
}

func generateSizedSurface(content string, multiline bool, font *ttf.Font) (*sdl.Surface, error) {
	var width, height int
	var err error
	if multiline && strings.Contains(content, "\n") {
		height = font.Height() * (strings.Count(content, "\n") + 1)
		lines := strings.Split(content, "\n")
		longest := 0
		for _, line := range lines {
			w, _, err := font.SizeUTF8(line)
			if err != nil {
				return nil, err
			}
			if w > longest {
				longest = w
			}
		}
		width = longest
	} else {
		width, height, err = font.SizeUTF8(content)
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("Generating a text surface of size", width, height)
	surf, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	surf.FillRect(nil, 0xFFFFFFFF)
	return surf, nil
}
