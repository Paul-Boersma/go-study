package phonebook

import "fmt"

type ContactList map[string]Contact

func (c ContactList) Add(contact Contact) {
	c[contact.Name] = Contact{contact.Name, contact.Phone}
}

func (c ContactList) List() {
	for _, contact := range c {
		fmt.Printf("%15v %15v\n", contact.Name, contact.Phone)
	}
}

func (c ContactList) SearchByName(name string) Contact {
	return c[name]
}

func (c ContactList) SearchByPhone(phone string) Contact {
	for _, contact := range c {
		if contact.Phone == phone {
			return contact
		}
	}
	return Contact{}
}
