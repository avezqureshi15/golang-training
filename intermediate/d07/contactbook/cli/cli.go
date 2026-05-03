package cli

import (
	"contactbook/models"
	"contactbook/storage"
	"fmt"
	"os"
	"strings"
)

func add(){
	var name, phone , email string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	contacts,_ := storage.LoadContacts()

	contacts = append(contacts, models.Contact{
		Name: name,
		Phone: phone,
		Email: email,
	})

	storage.SaveContacts(contacts)
	fmt.Println("Contact added successfully")
}

func list(){
	contacts,_ := storage.LoadContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts found")
		return
	}
	for i,c:= range contacts {
		fmt.Printf("%d. %s | %s | %s\n", i+1, c.Name, c.Phone, c.Email)
	}
}

func search(){
	if len(os.Args) < 3 {
		fmt.Println("Usage: search <name>")
		return
	}

	query := strings.ToLower(os.Args[2])
	contacts,_  := storage.LoadContacts()

	found := false
	for _,c := range contacts{
		if strings.Contains(strings.ToLower(c.Name),query){
			fmt.Printf("%s | %s | %s\n", c.Name, c.Phone, c.Email)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching contact found")
	}
}

func deleteContact(){
	if len(os.Args) < 3 {
		fmt.Println("Usage : delete <name>")
		return
	}

	query := strings.ToLower(os.Args[2])
	contacts,_ := storage.LoadContacts()

	var updated []models.Contact
	deleted := false

	for _,c := range contacts {
		if strings.ToLower(c.Name) != query{
			updated = append(updated, c)
		}else{
			deleted = true 
		}
	}

	storage.SaveContacts(updated)

	if deleted {
		fmt.Println("Contact added successfully")
	}else {
		fmt.Println("Contact not found")
	}

}

func Run() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: add | list | search | delete")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		add()
	case "list":
		list()
	case "search":
		search()
	case "delete":
		deleteContact()
	default:
		fmt.Println("Unknown command")
	}
}