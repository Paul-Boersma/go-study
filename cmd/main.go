package main

import (
	"go-study/internal/phonebook"
)

func main() {
	// Stating explicitly that this program creates a new phonebook each time it runs by providing the runner a contactList as input.
	var contactList phonebook.ContactList = make(map[string]phonebook.Contact)

	phonebook.RunProgram(contactList)
}
