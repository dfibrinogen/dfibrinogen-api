D'Fibrinogen API
================

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) installed.

```sh
$ go get -u github.com/dafian47/dfibrinogen-api
$ cd $GOPATH/src/github.com/dafian47/dfibrinogen-api
$ govendor sync
$ go run server.go
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [GoVendor](https://github.com/kardianos/govendor) if you are going to add any dependencies

## Built With

* [Gin Gonic](https://github.com/gin-gonic/gin)
* [Gorm](https://github.com/jinzhu/gorm)
* [PostgreSQL Driver](https://github.com/lib/pq)
* [Rest Secure](github.com/unrolled/secure)
* [Bcrypt](golang.org/x/crypto/bcrypt)

## ToDo

* [X] Rest for Login & Register
* [X] Rest for User, Category & Posts
* [ ] Rest for Likes & Comments
* [ ] Rest for Edit & Delete Posts
* [ ] Using JWT Token
* [ ] Documentation

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
