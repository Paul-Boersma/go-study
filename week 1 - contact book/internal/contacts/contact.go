package contacts

import "errors"

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func NewContact(name string, phone string) (Contact, error) {
	if name == "" {
		return Contact{}, errors.New("Name cannot be empty")
	}

	if phone == "" {
		return Contact{}, errors.New("Phone cannot be empty")
	}

	return Contact{name, phone}, nil
}
