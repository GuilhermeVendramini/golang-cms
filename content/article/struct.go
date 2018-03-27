package article

import "time"

// Article struct
type Article struct {
	Title   string
	Teaser  string
	Body    string
	Tags    string
	Author  string
	URL     string
	Changed time.Time
	Created time.Time
}
