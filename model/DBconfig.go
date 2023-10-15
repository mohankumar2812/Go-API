package model

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (user *User) Create() (*mongo.InsertOneResult, error) {
	userDB := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	emailCount, _ := userDB.CountDocuments(ctx, bson.M{"mail": user.Mail})

	defer cancel()

	if emailCount > 0 {
		return nil, errors.New("Email already exist")
	}

	result, err := userDB.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUser(mail string) (*User, error) {
	userDB := db.Collection("users")
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"mail": mail}

	err := userDB.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}
	
	return &user, nil

}

func UpdateUser(mail string, name string) (*mongo.UpdateResult, error) {
	userDB := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"mail": mail}
	updateData := bson.M{"$set": bson.M{"name": name}}

	opts := options.Update().SetUpsert(true)

	result, err := userDB.UpdateOne(ctx, filter, updateData, opts)

	if result.ModifiedCount == 0 {
		return nil, errors.New("given request not updated.")
	}

	if err != nil {
		return nil, errors.New("not updated.")
	}

	return result, nil

}

func DeleteUser(email string) (*mongo.DeleteResult, error) {

	userDB := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	filter := bson.M{"mail": email}

	result, err := userDB.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, nil

}
