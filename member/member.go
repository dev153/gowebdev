package member

import (
	"fmt"

	"github.com/dev153/gowebdev/person"
)

//Member structure defines a person with multiple skills.
type Member struct {
	person.Person
	Skills []string
}

//PrintSkills prints the skills of the member.
func (m Member) PrintSkills() {
	fmt.Println(m.Skills)
}

//PrintDetails prints the details of the "Member" structure type.
func (m Member) PrintDetails() {
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, skill := range m.Skills {
		fmt.Println("\t", skill)
	}
}
