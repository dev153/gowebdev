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
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]",
		p.Dob.String(), p.Email, p.Location)
}

//==============================================================================

type Admin struct {
	Person // type embedding for composition
	Roles  []string
}

//override PrintDetails from Person struct.
func (a Admin) PrintDetails() {
	//Call Person.PrintDetails
	a.Person.PrintDetails()
	fmt.Println("\nAdmin Roles:")
	for _, role := range a.Roles {
		fmt.Println("\t", role)
	}

}

//==============================================================================

type Member struct {
	Person // type embedding for composition
	Skills []string
}

//overrides PrintDetails from Person struct.
func (m Member) PrintDetails() {
	//Call Person.PrintDetails
	m.Person.PrintDetails()
	fmt.Println("\nSkills:")
	for _, skill := range m.Skills {
		fmt.Println("\t", skill)
	}
}

//==============================================================================

type Team struct {
	Name, Description string
	Users             []User
}

func (team Team) GetTeamDetails() {
	fmt.Printf("Team: %s - Description: %s", team.Name, team.Description)
	fmt.Println("Details of the team members:")
	for _, user := range team.Users {
		user.PrintName()
		user.PrintDetails()
	}
}

func main() {
	alex := Admin{
		Person{
			FirstName: "Alex",
			LastName:  "John",
			Dob:       time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			Email:     "alex.john@gmail.com",
			Location:  "New York",
		},
		[]string{"Manage Team", "Manage Tasks"},
	}

	shiju := Member{
		Person{
			FirstName: "Shiju",
			LastName:  "Varghese",
			Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			Email:     "shiju@gmail.com",
			Location:  "Kochi",
		},
		[]string{"Go", "Docker", "Kubernetes"},
	}

	chris := Member{
		Person{
			FirstName: "Chris",
			LastName:  "Martin",
			Dob:       time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
			Email:     "chris@gmail.com",
			Location:  "Santa Clara",
		},
		[]string{"Go", "Docker"},
	}

	team := Team{
		"Go",
		"Golang CoE",
		[]User{alex, shiju, chris},
	}

	team.GetTeamDetails()
}
