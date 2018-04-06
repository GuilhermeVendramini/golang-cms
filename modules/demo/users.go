package demo

import (
	"fmt"

	"github.com/GuilhermeVendramini/golang-cms/core/modules/users"
)

// User demo
func User() {
	var err error

	user := users.User{}
	user.Name = "admin"
	user.Email = "admin@admin.com"
	user.Admin = true
	user.Password, err = users.HashPassword("admin")
	if err != nil {
		panic(err)
	}

	_, err = users.Create(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Demo user as created:")
	fmt.Println("Name: admin, Email: admin@admin.com, Password: admin")
}
