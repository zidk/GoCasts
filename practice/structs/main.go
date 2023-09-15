package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

/*
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
*/

type person struct {
	firstName   string
	lastName    string
	contactInfo // if you don't want to use another name you can use this way
}

func main() {
	/*
		var alex person
		alex.firstName = "Alex"
		alex.lastName = "Anderson"
	*/
	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	alex.print()
	jim.print()
	//jimPointer := &jim
	jim.updateName("Paco")
	jim.print()

	name := "Bill"

	fmt.Println(*&name)
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}
