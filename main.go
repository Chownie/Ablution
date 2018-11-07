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

	container := ui.NewBaseContainer(80, 0, 100, 50, true)
	textfield := ui.NewTextField(0, 0, "Placeholder", true, `Font/UbuntuMono-R.ttf`, 14)
	textfield.SetRange(40, 100, 0, 50)
	button2 := ui.NewButton(0, 0, "Clear", `PNG\blue_button06.png`, `Font/UbuntuMono-R.ttf`, 12)
	button2.Register(ui.ButtonEvent, func(ev ui.EventData) {
		textfield.SetText("")
	})

	window.Add(button)
	container.Add(button2)
	container.Add(textfield)
	window.Add(container)
	fmt.Println(ui.Run(window))
}
