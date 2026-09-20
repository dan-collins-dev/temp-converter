package main

import (
	"fmt"
)

func main() {
	printTitle()
	printMainOptions()
	appLoop()
}

func toCelsius(fTemp float64) float64 {
	return (fTemp - 32) * (5.0 / 9.0)
}

func toFahrenheit(cTemp float64) float64 {
	return (cTemp * (9.0 / 5.0)) + 32
}

func printTitle() {
	fmt.Println("===============================")
	fmt.Println("| Temperature Converter in Go |")
	fmt.Println("===============================")
}

func printMainOptions() {
	fmt.Println("[1] Convert temperature from Fahrenheit to Celsius")
	fmt.Println("[2] Convert temperature from Celsius to Fahrenheit")
	fmt.Println("[3] Quit Program")
}

func appLoop() {
	for {
		var choice string

		fmt.Print("\nSelect an option: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		switch choice {
		case "1":
			fmt.Println("Fahrenheit Celsius")

		case "2":
			fmt.Println("Celsius to Fahrenheit")

		case "3":
			fmt.Println("Quitting application")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
