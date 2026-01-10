package phonebook

import "fmt"

func runProgram() {
	for {
		option, err := getOption()
		if err != nil {
			fmt.Println("error getting option by user")
			return
		}

		switch option {
		case "Add contact":
			contact, err := createContact()
			if err != nil {
				err = fmt.Errorf("Issue creating contact: %v", err)
				fmt.Println(err)
				return
			}
			fmt.Printf("Contact added: %v", contact)
		case "List contacts":
			listContacts()
		case "Search contacts":
			searchContacts()
		case "Exit":
			exitProgram()
		default:
			fmt.Println("Program exited unexpectedly")
			return
		}
	}
}
