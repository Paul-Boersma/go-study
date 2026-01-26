package persistance

import (
	"encoding/json"
	"go-study/internal/contacts"
	"os"
)

func LoadContactsFromFile(filename string) (contacts.ContactList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]contacts.Contact), nil
		}
		return nil, err
	} else if len(data) == 0 {
		return make(contacts.ContactList), nil
	}

	var contactList contacts.ContactList
	if err = json.Unmarshal(data, &contactList); err != nil {
		return nil, err
	}
	return contactList, nil
}

func SaveContactsToFile(filename string, contactList contacts.ContactList) error {
	data, err := json.Marshal(contactList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
