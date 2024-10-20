package users

import "github.com/andrei-kozel/owly-common/roles"

type User struct {
	Guid      string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Roles     []roles.Role
}
