package domain

type User struct {
	Guid      string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Roles     []Role
}
