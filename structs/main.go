package main

import "fmt"

func main() {
	john := person{"John", "Doe", contact{"john.doe@gmail.com", 12345}}

	jane := person{firstName: "Jane", lastName: "Doe", contact: contact{"jane.doe@gmail.com", 9876}}

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact = contact{"jane.doe@gmail.com", 45612}

	fmt.Println(john, jane, alex)
	fmt.Printf("%+v\n%+v\n%+v", john, jane, alex)
	john.print()
	jane.print()
	alex.print()

	alex.updateFirstName("Alexander")
	alex.print()

	alexPointer := &alex
	alexPointer.print()
	fmt.Println(*alexPointer)

	john.updateFirstName("Johnny")
	john.print()

	mySlice := []string{"How", "dead", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[1] = "deadish"
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email   string
	zipCode int
}
