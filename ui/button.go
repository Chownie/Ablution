package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Button struct {
	text            string
	parent          *Window
	font            *ttf.Font
	internalSurface *sdl.Surface
	slicedTexture   *sdl.Surface
	*EventManager
	*Transform
}

func NewButton(x int, y int, title string, imagePath string, fontPath string, fontsize int) *Button {
	tex, err := assetStore.GetImage(imagePath)
	if err != nil {
		panic(err)
	}
	surf, err := sdl.CreateRGBSurface(0, tex.W, tex.H, 32, 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	font, err := assetStore.GetFont(fontPath, fontsize)
	if err != nil {
		panic(err)
	}
	button := Button{title, nil, font, surf, tex, NewEventManager(), NewTransform(x, y, int(tex.W), int(tex.H), true)}
	return &button
}

func (b *Button) GetText() string {
	return b.text
}

func (b *Button) SetText(input string) {
	b.text = input
	b.SetDirty()
}

func (b *Button) SetParent(window *Window) {
	b.parent = window
	b.SetDirty()
}

// func (b *Button) Size() *sdl.Rect {
// 	return &sdl.Rect{int32(0), int32(0), b.w, b.h}
// }

// func (b *Button) Bounds() *sdl.Rect {
// 	return &sdl.Rect{b.x, b.y, b.w, b.h}
// }

// Redraw the internal surface with the new data
func (b *Button) SetDirty() {
	err := b.internalSurface.FillRect(&sdl.Rect{int32(0), int32(0),
		b.internalSurface.W, b.internalSurface.H}, 0xffffffff)
	if err != nil {
		panic(err)
	}
	b.slicedTexture.Blit(nil, b.internalSurface, nil)
	if len(b.text) > 0 {
		solid, err := b.font.RenderUTF8Solid(b.text, sdl.Color{255, 0, 0, 0})
		if err != nil {
			panic(err)
		}
		defer solid.Free()
		solid.Blit(nil, b.internalSurface, positionMiddle(solid, b.slicedTexture))
	}
	b.parent.SetDirty()
}

func (b *Button) Render() *sdl.Surface {
	return b.internalSurface
}
