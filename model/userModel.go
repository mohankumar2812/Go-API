package model

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       string `json: id`
	Name     string `json: name bson:"name,omitempty"`
	Mail     string `json: mail bson:"mail,omitempty"`
	PhNo     int    `json: phno bson:"phno,omitempty"`
	Password string `json: password bson:"password,omitempty"`
}

type JWTdesign struct{
	Email string
	jwt.RegisteredClaims
}

func (user *User) validate() (*User, error) {

	if strings.TrimSpace(user.Name) == "" {
		return nil, errors.New("name required")
	}

	if strings.TrimSpace(user.Mail) == "" {
		return nil, errors.New("mail required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("password requied")
	}

	return user, nil

}
