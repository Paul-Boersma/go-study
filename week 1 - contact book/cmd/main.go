package main

import (
	"fmt"
	"go-study/internal/cli"
	"go-study/internal/persistance"
)

func main() {
	const filename = "contactList.json"
	contactList, err := persistance.LoadContactsFromFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	cli.RunProgram(&contactList)

	if err = persistance.SaveContactsToFile(filename, contactList); err != nil {
		fmt.Println(err)
		return
	}
}
