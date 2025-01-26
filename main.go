package main

import "fmt"

type person struct {
	firstName	string
	lastName	string
	contactInfo
}

type contactInfo struct {
	email	string
	zip		int
}

func main() {
	paul := person {
		firstName: "Paul Michael",
		lastName: "Francisco",
		contactInfo: contactInfo{
			email: "pmofrancisco@yahoo.com",
			zip: 0,
		},
	}

	paul.updateName("Mek")
	paul.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}