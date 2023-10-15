package service

import (
	"errors"
	"example/model"
	"example/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {

	guid := xid.New()

	user.ID = guid.String()
	pass := utils.HashPassword(user.Password)
	user.Password = pass

	result, err := user.Create()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func LoginUser(user *model.User) (string, error) {

	login := user.Mail
	pass := user.Password

	result, err := model.GetUser(login)

	status := utils.CheckPassword(result.Password, pass)

	if !status {
		return "", errors.New("invalid password")
	}

	if err != nil {
		return "", errors.New("user not found")
	}

	loginToken := &model.JWTdesign{
		Email: user.Mail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Minute)),
			Subject:   string(user.Mail),
		},
	}

	token, err := CreateJWT(loginToken)

	if err != nil {
		return "", errors.New("JWT not generated")
	}


	return token, nil

}


