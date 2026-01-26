package cli

import (
	"fmt"
	"go-study/internal/contacts"
)

func RunProgram(contactList *contacts.ContactList) {
	commands := map[string]Cmd{
		"Add":    cmdAdd,
		"List":   cmdList,
		"Search": cmdSearch,
		"Edit":   cmdEdit,
		"Exit":   cmdExit,
	}

	for {
		option, err := provideMenu(commands)
		if err != nil {
			fmt.Println("error getting option by user")
			return
		}
		if option == nil {
			fmt.Println("Exiting Program")
			return
		}
		option(*contactList)
	}
}

func provideMenu(commands map[string]Cmd) (Cmd, error) {
	count := 1
	fmt.Println("Choose:")
	for i := range commands {
		fmt.Printf("%v.) %v\n", count, i)
		count++
	}

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return nil, err
	}

	cmd := commands[input]
	return cmd, nil
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

func cmdExit(contacts.ContactList) {
	fmt.Println("Program exited")
}
