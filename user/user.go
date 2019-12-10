package user

import "playground/doer"

// User is awesome too
type User struct {
	Doer doer.Doer
}

// Use is great
func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
