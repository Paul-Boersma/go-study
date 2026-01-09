package main

import (
	"fmt"
)

func main() {
	var conversionType int
	var temperature float64
	var unit string

	fmt.Println("Choose conversion type:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")

	n, err := fmt.Scanln(&conversionType)
	if err != nil || n != 1 {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Enter temperature: ")
	n, err = fmt.Scanln(&temperature)
	if err != nil || n != 1 {
		fmt.Println("Error reading input:", err)
	}

	switch conversionType {
	case 1:
		temperature = celciusToFahrenheit(temperature)
		unit = "Fahrenheit"
	case 2:
		temperature = fahrenheitToCelcius(temperature)
		unit = "Celsius"
	default:
		fmt.Println("Invalid conversion type selected.")
		return
	}

	fmt.Println("Converting...")
	fmt.Printf("Temperature: %v, %s", temperature, unit)
}

func celciusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
