package main

import (
	"fmt"

	"./ui"
)

func main() {
	ui.Init("assets")
	window := ui.NewWindow(400, 400, "Title")
	button := ui.NewButton(0, 0, "Submit", `PNG\blue_button06.png`, `Font/UbuntuMono-R.ttf`, 12)
	button.Register(ui.ButtonEvent, func(ev ui.EventData) {
		fmt.Println("main", window)
	})
	textfield := ui.NewTextField(96, 0, "Placeholder", true, `Font/UbuntuMono-R.ttf`, 14)
	textfield.SetRange(40, 200, 40, 100)
	button2 := ui.NewButton(48, 0, "Clear", `PNG\blue_button06.png`, `Font/UbuntuMono-R.ttf`, 12)
	button2.Register(ui.ButtonEvent, func(ev ui.EventData) {
		textfield.SetText("")
	})
	window.Add(button)
	window.Add(button2)
	window.Add(textfield)
	fmt.Println(ui.Run(window))
}
