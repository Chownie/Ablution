package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var assetStore *Assets
var selected Component

func Init(resourcePath string) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	assetStore = NewAssetManager(resourcePath)
}

func Run(window *Window) error {
	defer sdl.Quit()
	defer ttf.Quit()

	defer func() {
		for _, v := range assetStore.fontStore {
			v.Close()
		}
	}()

	win, err := sdl.CreateWindow(window.Title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		window.Width, window.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer win.Destroy()

	windowSurface, err := win.GetSurface()
	if err != nil {
		return err
	}

	for _, comp := range window.Components {
		comp.SetParent(window)
	}

	running := true
	for _, component := range window.Components {
		component.SetDirty()
	}
	window.SetDirty()
	for running {
		windowSurface.FillRect(&sdl.Rect{0, 0,
			windowSurface.W, windowSurface.H},
			0xffffff00)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseButtonEvent:
				uiEvent := EventData{ButtonEvent, sdl.Keysym{}, 0, t.Button, t.State, t.X, t.Y, "", 0, 0}
				for _, component := range window.Components {
					if component.Contains(t.X, t.Y) {
						if selected != component {
							if selected != nil {
								selected.Trigger(DeselectEvent, uiEvent)
							}
							component.Trigger(SelectEvent, uiEvent)
							selected = component
						}
						component.Trigger(ButtonEvent, uiEvent)
					}
				}
			case *sdl.MouseMotionEvent:
				uiEvent := EventData{HoverEvent, sdl.Keysym{}, 0, 0, 0, t.X, t.Y, "", 0, 0}
				for _, component := range window.Components {
					if component.Contains(t.X, t.Y) {
						component.Trigger(HoverEvent, uiEvent)
					}
				}
			case *sdl.KeyboardEvent:
				if selected == nil {
					break
				}
				uiEvent := EventData{KeyEvent, t.Keysym, t.State, 0, 0, 0, 0, "", 0, 0}
				selected.Trigger(KeyEvent, uiEvent)
			case *sdl.TextInputEvent:
				if selected == nil {
					break
				}
				uiEvent := EventData{TextEvent, sdl.Keysym{}, 0, 0, 0, 0, 0, string(t.Text[:]), 0, 0}
				selected.Trigger(TextEvent, uiEvent)
			case *sdl.TextEditingEvent:
				if selected == nil {
					break
				}
				uiEvent := EventData{TextEvent, sdl.Keysym{}, 0, 0, 0, 0, 0, string(t.Text[:]), t.Start, t.Length}
				selected.Trigger(TextEditEvent, uiEvent)
			}

		}
		window.Render().Blit(&sdl.Rect{0, 0, windowSurface.W, windowSurface.H},
			windowSurface, &sdl.Rect{0, 0, windowSurface.W, windowSurface.H})
		win.UpdateSurface()
	}

	return nil
}
