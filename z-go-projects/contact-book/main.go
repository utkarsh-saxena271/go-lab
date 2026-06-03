package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

var contacts []Contact

func addContact() {
	var name string
	var phone string

	fmt.Print("Enter Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Phone: ")
	fmt.Scan(&phone)

	contact := Contact{
		Name:  name,
		Phone: phone,
	}

	contacts = append(contacts, contact)

	fmt.Println("Contact Added!")
}

func viewContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\nContacts:")
	for i, contact := range contacts {
		fmt.Printf("%d. %s - %s\n", i+1, contact.Name, contact.Phone)
	}
}

func searchContact() {
	var name string

	fmt.Print("Enter Name: ")
	fmt.Scan(&name)

	for _, contact := range contacts {
		if contact.Name == name {
			fmt.Println("Found:")
			fmt.Printf("%s - %s\n", contact.Name, contact.Phone)
			return
		}
	}

	fmt.Println("Contact not found.")
}

func deleteContact() {
	var name string

	fmt.Print("Enter Name: ")
	fmt.Scan(&name)

	for i, contact := range contacts {
		if contact.Name == name {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Println("Contact deleted.")
			return
		}
	}

	fmt.Println("Contact not found.")
}

func main() {
	for {
		fmt.Println("\n===== Contact Book =====")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addContact()
		case 2:
			viewContacts()
		case 3:
			searchContact()
		case 4:
			deleteContact()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}