package main

import (
	"encoding/json"
	"fmt"
	"go-study/internal/phonebook"
	"os"
)

func main() {
	// Stating explicitly that this program creates a new phonebook each time it runs by providing the runner a contactList as input.

	// Save to File
	// Load from File
	// Convert to JSON
	const filename = "contactList.json"
	contactList, err := LoadFromFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	phonebook.RunProgram(&contactList)

	if err = SaveToFile(filename, contactList); err != nil {
		fmt.Println(err)
		return
	}

}

func LoadFromFile(filename string) (phonebook.ContactList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]phonebook.Contact), nil
		}
		return nil, err
	} else if len(data) == 0 {
		return make(phonebook.ContactList), nil
	}

	var contactList phonebook.ContactList
	if err = json.Unmarshal(data, &contactList); err != nil {
		return nil, err
	}
	return contactList, nil
}

func SaveToFile(filename string, contactList phonebook.ContactList) error {
	data, err := json.Marshal(contactList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
