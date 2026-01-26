package cli

import (
	"fmt"
	"go-study/internal/contacts"
)

type Cmd func(contacts.ContactList)

func cmdAdd(c contacts.ContactList) {
	name, err := getName()
	if err != nil {
		fmt.Println(err)
		return
	}

	phone, err := getPhoneNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	contact, err := contacts.NewContact(name, phone)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Add(contact)
}

func cmdList(c contacts.ContactList) {
	c.List()
}

func cmdSearch(c contacts.ContactList) {
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
		contact := c.SearchByName(name)
		fmt.Printf("%15v %15v\n", contact.Name, contact.Phone)
	case 2:
		phone, err := getPhoneNumber()
		if err != nil {
			fmt.Println(err)
		}
		contact := c.SearchByPhone(phone)
		fmt.Printf("%15v %15v\n", contact.Name, contact.Phone)
	}
}

func cmdEdit(c contacts.ContactList) {

}
