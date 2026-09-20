package main

import "fmt"

func main() {
	todaysTemp := toCelsius(91.4)
	tomorrowsTemp := toFahrenheit(0)

	fmt.Println(todaysTemp, tomorrowsTemp)
}

func toCelsius(fTemp float64) float64 {
	return (fTemp - 32) * (5.0 / 9.0)
}

func toFahrenheit(cTemp float64) float64 {
	return (cTemp * (9.0 / 5.0)) + 32
}
