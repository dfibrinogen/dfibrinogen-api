D'Fibrinogen API
================

## Getting Start

You can open this project with Sublime Text, Atom or GoLand

## Prerequisites

What things you need to install the software and install them :

* [Go](http://golang.org/doc/install)
* [PostgreSQL](https://www.postgresql.org/)
* [Dep](https://github.com/golang/dep)

If you using Mac OS, you could install this with [Homebrew](homebrew.sh)

```sh
$ brew install go
$ brew install postgresql
$ brew install dep
```

## Testing

Just run this command

```sh
$ make test
```

Generate mock using Mockery

```sh
$ $GOPATH/bin/mockery -name=HelloInterface -inpkg
```

## ToDo

* DB Migration
* Test Connection Database
* Using Token

## Built With

* [Echo](https://github.com/labstack/echo)
* [Gorm](https://github.com/jinzhu/gorm)
* [Bcrypt](golang.org/x/crypto/bcrypt)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
