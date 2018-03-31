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
func Update(user User, ID string) (User, error) {
	err := Users.Update(bson.M{"_id": bson.ObjectIdHex(ID)}, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Remove user
func Remove(ID string) error {
	err := Users.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
	if err != nil {
		return errors.New("500 internal server error")
	}
	return nil
}

// GetbyID return one user by ID
func GetbyID(ID string) (User, error) {
	user := User{}
	err := Users.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&user)
	return user, err
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
