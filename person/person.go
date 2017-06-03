package person

import (
	"fmt"
	"time"
)

//Person structure contains a person's basic information data.
type Person struct {
	FirstName string
	LastName  string
	DoB       time.Time
	Email     string
	Location  string
}

//PrintName prints the first and last name of the person.
func (p Person) PrintName() {
	fmt.Println(p.FirstName, p.LastName)
}

//PrintDetails print the details of the person information.
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.DoB.String(), p.Email, p.Location)
}

//ChangeLocation changes the "Location" structure field.
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}
