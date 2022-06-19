package entity

import (
	"time"

	"github.com/chkilel/fiberent/ent"
	"golang.org/x/crypto/bcrypt"
)

// User is the model entity for the User schema.
type User struct {
	ent.User
}

// NewUser create a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		ent.User{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			CreatedAt: time.Now(),
		},
	}

	pwd, err := u.GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	return u, nil
}

// ValidatePassword validate user password
func ValidatePassword(u *User, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

// generatePassword generate password
func (u *User) GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
