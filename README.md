
# gh0st
Cross-platform, lightweight and simple reporting framework.

It is a simple **RESTful API**, that can be consumed by any
HTTP client. This makes easy to integrate external tools with it.

## Features

* Simple RESTful API
* Client SDK for **Go** (see `sdk/`)
* Cross-platform and lightweight
* Easy integration with other tools (see `external/`)

## Prerequisites
You must have some things before you can run this:

* [Go](https://golang.org/) > 1.8
* [GNU Make](https://www.gnu.org/software/make/)

## Getting Started
Run `make build` to get the binaries in the directory `./dist/`.

Run the `schema.sql` file to setup a PostgreSQL database. Copy the `.gh0std.toml`
file to `$HOME`, and modify your database credentials on it. Finally, run then `gh0std`
binary to start the server.

## Built With

* [Go](https://golang.org/) - Programming Language
* [dep](https://github.com/golang/dep) - Vendoring Tool
* [PostgreSQL](https://www.postgresql.org/) - Data Storage
* [gin](https://github.com/gin-gonic/gin) - Web Framework
* [go-pg](https://github.com/go-pg/pg) - PostgreSQL ORM

## Versioning
I use [SemVer](http://semver.org/) for versioning.
For the versions available, see the tags on this repository.

## Authors
* **Gustavo Paulo** - [gpaulo00](https://github.com/gpaulo00)
