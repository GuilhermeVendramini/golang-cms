
# Golang CMS

Build fast and flexible websites in Golang.

# How to use

Change your mongodb credentials and database name in "config/db.go".
Uncomment the line "demo.User()" inside of main.go.

"demo.User()" will gerenate a demo user:

> **Name:** admin

> **Email:** admin@admin.com

> **Password:** admin

Execute:

```sh
$ go run main.go
```

After running for the first time, comment "demo.User()" otherwise that will generate a new user again.

**Enjoy!**

# Packages

[gopkg.in/mgo.v2](https://gopkg.in/mgo.v2)

[github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

[github.com/gorilla/securecookie](https://github.com/gorilla/securecookie)

[golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

[gopkg.in/gomail.v2](https://gopkg.in/gomail.v2)

[github.com/go-redis/redis](https://github.com/go-redis/redis)