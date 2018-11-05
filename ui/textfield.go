package ui

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextField struct {
	text            string
	parent          *Window
	internalSurface *sdl.Surface
	Multiline       bool
	font            *ttf.Font
	selected        bool
	*EventManager
	*Transform
}

func NewTextField(x int, y int, placeholder string, multiline bool, fontface string, fontsize int) *TextField {
	font, err := assetStore.GetFont(fontface, fontsize)
	if err != nil {
		panic(err)
	}

	surf, err := generateSizedSurface(placeholder, multiline, font)
	if err != nil {
		panic(err)
	}

	TextField := TextField{placeholder, nil, surf, multiline, font, false, NewEventManager(), NewTransform(x, y, int(surf.W), int(surf.H), false)}

	TextField.Register(SelectEvent, func(ev EventData) {
		TextField.selected = true
		sdl.StartTextInput()
	})
	TextField.Register(TextEvent, func(ev EventData) {
		if TextField.selected == true {
			TextField.Append(string(ev.Text[0]))
		}
		/* && ev.KeyState == 0 {
			//fmt.Println(ev.KeyState)
			switch ev.Key.Sym {
			case sdl.K_BACKSPACE:
				TextField.Remove(1)
			case sdl.K_RETURN:
				fallthrough
			case sdl.K_RETURN2:
				TextField.text += "\n"
			default:
				TextField.Append(sdl.GetKeyName(ev.Key.Sym))
			}
		}*/
	})
	TextField.Register(KeyEvent, func(ev EventData) {
		if TextField.selected && ev.KeyState == 1 {
			switch ev.Key.Sym {
			case sdl.K_BACKSPACE:
				TextField.Remove(1)
			case sdl.K_RETURN:
				fallthrough
			case sdl.K_RETURN2:
				TextField.text += "\n"
			}
		}
	})
	TextField.Register(DeselectEvent, func(ev EventData) {
		TextField.selected = false
		sdl.StopTextInput()
	})

	return &TextField
}

func (tx *TextField) Remove(length int) {
	if len(tx.text) > length-1 {
		tx.text = tx.text[:len(tx.text)-length]
	}
	tx.SetDirty()
}

func (tx *TextField) Append(input string) {
	tx.text += input
	tx.SetDirty()
}

func (tx *TextField) SetText(input string) {
	tx.text = input
	tx.SetDirty()
}

func (tx *TextField) GetText() string {
	return tx.text
}

func (tx *TextField) Size() *sdl.Rect {
	return &sdl.Rect{int32(0), int32(0), tx.internalSurface.W, tx.internalSurface.H}
}

func (tx *TextField) Bounds() *sdl.Rect {
	return &sdl.Rect{tx.X, tx.Y, tx.internalSurface.W, tx.internalSurface.H}
}

func (tx *TextField) SetParent(win *Window) {
	tx.parent = win
}

func (tx *TextField) SetDirty() {
	var err error
	tx.internalSurface, err = generateSizedSurface(tx.text, tx.Multiline, tx.font)
	if err != nil {
		panic(err)
	}
	if tx.Multiline && strings.Contains(tx.text, "\n") {
		lines := strings.Split(tx.text, "\n")
		for offset, line := range lines {
			solid, err := tx.font.RenderUTF8Solid(line, sdl.Color{0, 0, 0, 255})
			if err != nil {
				break
			}
			defer solid.Free()
			yOffset := offset * tx.font.Height()
			solid.Blit(nil, tx.internalSurface, &sdl.Rect{int32(0), int32(yOffset), solid.W, solid.H})
		}
	} else {
		solid, err := tx.font.RenderUTF8Solid(tx.text, sdl.Color{0, 0, 0, 255})
		if err != nil {
			fmt.Println(err)
		}
		defer solid.Free()
		solid.Blit(nil, tx.internalSurface, nil)
	}

	tx.parent.SetDirty()
}

func (tx *TextField) Render() *sdl.Surface {
	return tx.internalSurface
}
