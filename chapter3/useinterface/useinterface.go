package main

import (
	"fmt"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n",
		p.Dob.String(), p.Email, p.Location)
}

func main() {
	alex := Person{
		FirstName: "Alex",
		LastName:  "John",
		Dob:       time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
		Email:     "alex.john@gmail.com",
		Location:  "New York",
	}

	shiju := Person{
		FirstName: "Shiju",
		LastName:  "Varghese",
		Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		Email:     "shiju@gmail.com",
		Location:  "Kochi",
	}

	users := []Person{alex, shiju}
	for _, aUser := range users {
		aUser.PrintName()
		aUser.PrintDetails()
	}
}
