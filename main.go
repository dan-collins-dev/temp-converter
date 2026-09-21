package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func main() {
	tempInput := tview.NewInputField().
		SetLabel("Input the temperature you would like to convert: ").SetFieldWidth(10)

	dropdown := tview.NewDropDown().
		SetLabel("Select the unit you want to convert to: ").
		SetOptions([]string{"Celsius", "Fahrenheit"}, nil)

	result := tview.NewTextView().SetText("Result: ")
	app := tview.NewApplication()

	form := tview.NewForm().
		AddFormItem(tempInput).
		AddFormItem(dropdown).
		AddFormItem(result).
		AddButton("Submit", func() {
			value, err := strconv.ParseFloat(tempInput.GetText(), 64)
			if err != nil {
				result.SetText("Result: Please enter a valid number")
				return
			}

			selectedIndex, selected := dropdown.GetCurrentOption()

			var output string

			if selectedIndex == 0 {
				output = fmt.Sprintf("Result: %.1f°F = %.1f°C", value, toCelsius(value))
			} else if selectedIndex == 1 {
				output = fmt.Sprintf("Result: %.1f°C = %.1f°F", value, toFahrenheit(value))
			} else {
				result.SetText("A unit must be selected")
				return
			}
			_ = selected

			result.SetText(output)
			tempInput.SetText("")
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	app.SetRoot(form, true)

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func toCelsius(fTemp float64) float64 {
	return (fTemp - 32) * (5.0 / 9.0)
}

func toFahrenheit(cTemp float64) float64 {
	return (cTemp * (9.0 / 5.0)) + 32
}
