package user

type Email string

type User struct {
	UID      UID
	Email    Email
	Password HashedPassword
}
