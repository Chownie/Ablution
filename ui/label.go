package ui

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Label struct {
	text            string
	parent          *Window
	internalSurface *sdl.Surface
	Multiline       bool
	font            *ttf.Font
	*EventManager
	*Transform
}

func NewLabel(x int, y int, content string, multiline bool, fontface string, fontsize int) *Label {
	font, err := assetStore.GetFont(fontface, fontsize)
	if err != nil {
		panic(err)
	}

	surf, err := generateSizedSurface(content, multiline, font)
	if err != nil {
		panic(err)
	}

	label := Label{content, nil, surf, multiline, font, NewEventManager(), new(Transform)}
	return &label
}

func (l *Label) GetText() string {
	return l.text
}

func (l *Label) SetText(input string) {
	l.text = input
	l.SetDirty()
}

func (l *Label) SetDirty() {
	var err error
	l.internalSurface, err = generateSizedSurface(l.text, l.Multiline, l.font)
	if err != nil {
		panic(err)
	}
	if l.Multiline && strings.Contains(l.text, "\n") {
		lines := strings.Split(l.text, "\n")
		for offset, line := range lines {
			solid, err := l.font.RenderUTF8Solid(line, sdl.Color{0, 0, 0, 255})
			if err != nil {
				fmt.Println(err)
			}
			defer solid.Free()
			yOffset := offset * l.font.Height()
			solid.Blit(nil, l.internalSurface, &sdl.Rect{int32(0), int32(yOffset), solid.W, solid.H})
		}
	} else {
		solid, err := l.font.RenderUTF8Solid(l.text, sdl.Color{0, 0, 0, 255})
		if err != nil {
			fmt.Println(err)
		}
		defer solid.Free()
		solid.Blit(nil, l.internalSurface, nil)
	}
	l.parent.SetDirty()
}

func (l *Label) Render() *sdl.Surface {
	return l.internalSurface
}

func (l *Label) SetParent(window *Window) {
	l.parent = window
}
