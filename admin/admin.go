package admin

import (
	"fmt"

	personPkg "github.com/dev153/gowebdev/person"
)

//Admin structure defines a person with multiple roles.
type Admin struct {
	Person personPkg.Person
	Roles  []string
}

//PrintRoles prints the roles of the admin.
func (a Admin) PrintRoles() {
	fmt.Println(a.Roles)
}
