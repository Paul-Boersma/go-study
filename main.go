package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Exercise day 4

	// print usefull statistics about the numbers
	var data data

	input, err := askUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	data.Count = len(input)
	for _, value := range input {
		data.Min = math.Min(data.Min, value)
		data.Max = math.Max(data.Max, value)
		data.Sum += value
	}
	data.Avg = data.Sum / float64(data.Count)

	fmt.Printf(`Data:
Count: %v
Min: %v
Max: %v
Sum: %v
Avg: %.2f
`, data.Count, data.Min, data.Max, data.Sum, data.Avg)
}

func askUserInput() ([]float64, error) {
	var userInput []float64
	var input string

	fmt.Println("Enter digits. When ready type 'done'")
	for {
		fmt.Scanf("%v\n", &input)
		if input == "done" {
			return userInput, nil
		}

		formattedInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("input not valid, please enter digits or type 'done' to exit")
			continue
		}

		userInput = append(userInput, formattedInput)
	}
}

type data struct {
	Count int
	Sum   float64
	Avg   float64
	Min   float64
	Max   float64
}
