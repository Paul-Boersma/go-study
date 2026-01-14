package phonebook

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func createContact() (Contact, error) {
	name, err := getName()
	if err != nil {
		return Contact{}, err
	}

	phone, err := getPhoneNumber()
	if err != nil {
		return Contact{}, err
	}
	return Contact{name, phone}, nil
}

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

	// if err := validatePhoneNumber(phoneNumber); err != nil {
	// 	return "", err
	// }

	return phoneNumber, nil
}

func validatePhoneNumber(phoneNumber string) error {
	if _, err := fmt.Sscanf(phoneNumber, "%2d-%6d\n"); err != nil {
		return err
	}
	return nil
}
