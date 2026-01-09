package main

import "fmt"

func main() {
	operation, err := requestOperation()
	if err != nil {
		return
	}

	input1, input2, err := readValuesInput()
	if err != nil {
		return
	}

	result, err := performOperation(operation, input1, input2)
	if err != nil {
		fmt.Printf("Error performing calculation: %v", err)
		return
	}

	fmt.Printf("Calculation result: %v", result)
}

func requestOperation() (input int, err error) {
	fmt.Println(`Choose operation:
	1) Add
	2) Subtract
	3) Multiply
	4) Divide`)

	_, err = fmt.Scanf("%d\n", &input)
	if err != nil || input <= 0 || input > 4 {
		return 0, err
	}
	return input, nil
}

func readValuesInput() (input1 float64, input2 float64, err error) {
	fmt.Println("Provide numbers you'd like to use in the following format: %f %f")
	_, err = fmt.Scanf("%f %f", &input1, &input2)
	if err != nil {
		return
	}

	return input1, input2, nil
}

func performOperation(operation int, input1 float64, input2 float64) (float64, error) {
	switch operation {
	case 1:
		return add(input1, input2), nil
	case 2:
		return sub(input1, input2), nil
	case 3:
		return multiply(input1, input2), nil
	case 4:
		return divide(input1, input2)
	default:
		return 0, fmt.Errorf("Invalid operation choice")
	}
}

func add(a float64, b float64) float64 {
	return a + b
}

func sub(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

// How to optionally return?
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0.")
	}
	return (a / b), nil
}
