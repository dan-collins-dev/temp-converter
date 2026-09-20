package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	printTitle()
	appLoop()
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
		printMainOptions()
		choice, err := getChoice()

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			tempInput, err := getTemperature()

			if err != nil {
				fmt.Println(err)
				continue
			}

			temperature := toCelsius(tempInput)
			fmt.Printf("\nThe entered temperature in degrees celsius is %.1f\u00b0C\n\n", temperature)

		case 2:
			tempInput, err := getTemperature()

			if err != nil {
				fmt.Println(err)
				continue
			}

			temperature := toFahrenheit(tempInput)
			fmt.Printf("\nThe entered temperature in degrees fahrenheit is %.1f\u00b0C\n\n", temperature)

		case 3:
			fmt.Println("\nQuitting application")
			return

		default:
			fmt.Print("Invalid option.\n\n")
		}
	}
}

func getChoice() (int, error) {
	var input string

	fmt.Print("\nSelect an option: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		return 0, errors.New("Error reading input")
	}

	num, err := strconv.Atoi(input)

	if err != nil {
		return 0, errors.New("\nEntered input cannot be converted to an integer.\n")
	}

	return num, nil
}

func getTemperature() (float64, error) {
	var input string

	fmt.Print("Enter the temperature to convert: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		return 0, errors.New("Error reading input")
	}

	temp, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, errors.New("\nEntered input cannot be converted to a float.\n")
	}

	return temp, nil
}

func toCelsius(fTemp float64) float64 {
	return (fTemp - 32) * (5.0 / 9.0)
}

func toFahrenheit(cTemp float64) float64 {
	return (cTemp * (9.0 / 5.0)) + 32
}
