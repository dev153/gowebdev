package main

import (
	"fmt"
	"time"

	"github.com/dev153/gowebdev/admin"
	"github.com/dev153/gowebdev/member"
	"github.com/dev153/gowebdev/person"
)

func main() {
	p := person.Person{
		FirstName: "Shiju",
		LastName:  "Varghese",
		DoB:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		Email:     "shiju@gmail.com",
		Location:  "Kochi",
	}
	/* Create a pointer because the signature of "ChangeLocation"
	   has as a receiver a pointer to the Person structure. */
	ptr := &p
	p.PrintName()
	p.PrintDetails()
	ptr.ChangeLocation("New Delhi")
	p.PrintDetails()

	panos := admin.Admin{
		Person: p,
		Roles:  []string{"Manage Team", "Manage Tasks"},
	}

	nick := member.Member{
		Person: p,
		Skills: []string{"Go", "Kubernetes", "Docker"},
	}

	fmt.Println(panos)
	fmt.Println(nick)
	nick.PrintSkills()
	panos.PrintRoles()
	nick.PrintName()
	nick.PrintDetails()
}
