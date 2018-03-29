package users

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Create a new User
func Create(user User) (User, error) {
	err := Users.Insert(user)
	if err != nil {
		return user, errors.New("internal server error" + err.Error())
	}
	return user, nil
}

// Update user
func Update(user User, currentEmail string) (User, error) {
	err := Users.Update(bson.M{"url": currentEmail}, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetAll return all users
func GetAll() ([]User, error) {
	users := []User{}
	err := Users.Find(bson.M{}).Sort("-_id").All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
