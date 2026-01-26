package cli

import "fmt"

func getName() (name string, err error) {
	fmt.Println("Please insert your name:")

	_, err = fmt.Scanln(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func getPhoneNumber() (phoneNumber string, err error) {
	fmt.Println("Please insert your phone number:")

	fmt.Scanln(&phoneNumber)

	return phoneNumber, nil
}
