package phonebook

import "fmt"

func RunProgram(contactList ContactList) {
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
				fmt.Println(err)
				return
			}
			contactList.Add(contact)
		case "List contacts":
			contactList.List()
		case "Search contacts":
			searchOption, err := getSearchOption()
			if err != nil {
				fmt.Println(err)
				return
			}
			switch searchOption {
			case 1:
				name, err := getName()
				if err != nil {
					fmt.Println(err)
				}
				contact := contactList.SearchByName(name)
				fmt.Printf("%15v %15v\n", contact.Name, contact.Phone)
			case 2:
				phone, err := getPhoneNumber()
				if err != nil {
					fmt.Println(err)
				}
				contact := contactList.SearchByPhone(phone)
				fmt.Printf("%15v %15v\n", contact.Name, contact.Phone)
			}
		case "Exit":
			exitProgram()
		default:
			fmt.Println("Program exited unexpectedly")
			return
		}
	}
}

func getOption() (option string, err error) {
	fmt.Println(`Choose:
1.) Add contact
2.) List contacts
3.) Search contacts
4.) Exit program`)

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
		option = "Search contacts"
	case 4:
		option = "Exit Program"
	default:

	}

	return option, nil

}

func getSearchOption() (int, error) {
	fmt.Println(`How do you want to search the contact list?
1.) By name
2.) By phone`)

	var option int
	if _, err := fmt.Scanln(&option); err != nil {
		return 0, err
	}
	return option, nil
}

func exitProgram() {
	fmt.Println("Program exited")
}
