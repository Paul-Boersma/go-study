package phonebook

import "fmt"

type ContactList struct {
	Contacts []Contact
}

func (c *ContactList) Search() {
	for _, contact := range c.Contacts {
		fmt.Scanln("%v", contact.Name)
	}
}

func (c *ContactList) List() {
	for _, contact := range c.Contacts {
		fmt.Printf("%15v %15v", contact.Name, contact.Phone)
	}
}

func (c *ContactList) Add(contact Contact) {
	c.Contacts = append(c.Contacts, contact)
}

func (c *ContactList) CreateContact() {
	name, err := getName()
	if err != nil {
		return Contact{}, err
	}

	phoneNumber, err := getPhoneNumber()
	if err != nil {
		return Contact{}, err
	}

	c.Add(Contact{})

	return Contact{
		Name:  name,
		Phone: phoneNumber,
	}, nil
}
