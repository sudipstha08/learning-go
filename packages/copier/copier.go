// https://github.com/lunixbochs/go-copier
package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Role string
	Age  int32
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name       string
	Age        int32
	DoubleAge  int32
	EmployeeId int64
	SuperRule  string
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func main() {
	user := User{Name: "John Cena", Age: 18, Role: "Admin"}
	employee := Employee{}

	// Copy struct to struct
	copier.Copy(&employee, &user)
	fmt.Println("employee", employee)

	// Copy slice to slice
	users := []User{{Name: "Micheal", Age: 18, Role: "Admin"}, {Name: "John", Age: 30, Role: "Developer"}}
	employees := []Employee{} // slices of structs
	copier.Copy(&employees, &users)
	fmt.Println("employees: ", employees)
}

// FEATURES
// -> Copy field to field if name exactly match
// -> Copy from method to field if method name and field name exactly match
// -> Copy from field to method if field name and method name exactly match
// -> Copy slice to slice
// -> Copy struct to slice
