package main

import "fmt"

func main() {

	runProgram()

	// Search for contact by name

	// Exit the program

}

func getOption() (option string, err error) {
	fmt.Println(`Choose:
1.) Add contact
2.) List contacts
3.) Exit program`)

	var input int
	if _, err := fmt.Scanln(&input); err != nil {
		return "", err
	}

	switch input {
	case 1:
		option = "Add contact"
	case 2:
		option = "List contacts"
	case 3:
		option = "Exit Program"
	}

	return option, nil

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

	if err := validatePhoneNumber(phoneNumber); err != nil {
		return "", err
	}

	return phoneNumber, nil
}

func validatePhoneNumber(phoneNumber string) error {
	if _, err := fmt.Sscanf(phoneNumber, "%2d-%6d\n"); err != nil {
		return err
	}
	return nil
}

func listContacts() {

}

func searchContacts(contactList []Contact) (Contact, error) {
	fmt.Println("How do you want to search?")

	var option string

	if _, err := fmt.Scanln(&option); err != nil {
		return Contact{}, err
	}

	switch option {
	case "by Name":
		contact, err := searchContactByName()
	case "by Phone Number":
		contact, err := searchByPhoneNumber()
	}
	return contact, nil
}

func searchByName() (Contact, error) {
	// get input name
	// search contact list
	return Contact{}, nil
}

func searchByPhoneNumber() (Contact, error) {
	// get input phone number
	// search contactlist

	return Contact{}, nil
}

func exitProgram() {
	fmt.Println("Program exited")
}
