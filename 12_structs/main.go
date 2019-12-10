package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	// return fmt.Sprintln("Hello, my name is", person.firstName, person.lastName, "and I am", person.age)
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) newBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	switch person.gender {
	case "f":
		person.lastName = spouseLastName
	default:
		return
	}
}

func main() {
	// Initialize a person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       25}

	// Alternative
	// person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.newBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
