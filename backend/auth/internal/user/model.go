package user

type ID string

type Email string

type Password string

type User struct {
	ID       ID
	Email    Email
	Password Password
}
