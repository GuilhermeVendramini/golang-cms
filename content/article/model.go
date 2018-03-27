package article

import (
	"errors"
)

// Create a new Article
func Create(item Article) (Article, error) {
	err := Articles.Insert(item)
	if err != nil {
		return item, errors.New("internal server error" + err.Error())
	}
	return item, nil
}
