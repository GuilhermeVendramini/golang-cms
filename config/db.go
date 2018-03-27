package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// DB database
var DB *mgo.Database

func init() {
	// Your mongodb connection
	s, err := mgo.Dial("mongodb://localhost/golangcms")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	// Your database name
	DB = s.DB("golangcms")
	fmt.Println("You connected to your mongo database.")
}
