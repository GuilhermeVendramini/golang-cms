package users

import (
	"github.com/GuilhermeVendramini/golang-cms/config"
	"gopkg.in/mgo.v2"
)

// Users collection
var Users *mgo.Collection

func init() {
	Users = config.DB.C("users")
}
