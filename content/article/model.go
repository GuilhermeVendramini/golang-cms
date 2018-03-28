package article

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Create a new Article
func Create(item Article) (Article, error) {
	err := Articles.Insert(item)
	if err != nil {
		return item, errors.New("internal server error" + err.Error())
	}
	return item, nil
}

// Get return one article
func Get(URL string) (Article, error) {
	item := Article{}
	err := Articles.Find(bson.M{"url": URL}).One(&item)
	return item, err
}

// Remove article
func Remove(URL string) error {
	err := Articles.Remove(bson.M{"url": URL})
	if err != nil {
		return errors.New("500 internal server error")
	}
	return nil
}

// Update article
func Update(item Article, currentURL string) (Article, error) {
	err := Articles.Update(bson.M{"url": currentURL}, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// GetAll return all articles
func GetAll() ([]Article, error) {
	items := []Article{}
	err := Articles.Find(bson.M{}).Sort("-_id").All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
