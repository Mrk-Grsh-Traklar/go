package main

import (
	// "fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// кодовое название "Ты хули жмешь, пидор <3"
// func main() {
// 	a := app.New()
// 	w := a.NewWindow("bebra")
// 	label := widget.NewLabel("Hello World!")
// 	label2 := widget.NewLabel("Fuck off <3")
// 	btn := widget.NewButton("Press me :)", func() {
// 		fmt.Println("Ты хули жмешь, пидор <3")
// 	})

// 	w.SetContent(container.NewVBox(
// 		label,
// 		label2,
// 		btn,
// 	))
// 	w.ShowAndRun()
// }

func main() {
	a := app.New()
	w := a.NewWindow("stupid culculator")
	label1 := widget.NewLabel("first number")
	entry := widget.NewEntry()

	label2 := widget.NewLabel("second number")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("answer")

	btn := widget.NewButton("count", func() {
		n1, err := strconv.ParseFloat(entry.Text, 64)
		n2, er := strconv.ParseFloat(entry2.Text, 64)

		if err != nil || er != nil {
			answer.SetText("Error")
		}

		sum := n1 + n2
		sum := n1 - n2
		sum := n1 * n2
		sum := n1 + n2
	})

	w.SetContent(container.NewVBox(
		label1,
		entry,
		label2,
		entry2,
		answer,
		btn,
	))
	w.ShowAndRun()
}
