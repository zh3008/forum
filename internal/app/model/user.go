package model

import "golang.org/x/crypto/bcrypt"

//User ..
type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"omitempty"`
}

//NewUser ..
func NewUser() *User {
	return &User{}
}

//BeforeCreate ..
func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := EncryptString(u.Password)
		if err != nil {
			return err
		}

		u.Password = enc
	}

	return nil
}

//EncryptString ..
func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
