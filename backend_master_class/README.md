# Backend Master Class

Code and notes from studying [Backend Master Class \[Golang + Postgres + Kubernetes + gRPC\]](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/)

The course walks to a Simple bank project

## Section 1: Working with database [Postgres + SQLC]

### Lecture 1: Design DB schema

Cloned project at [github cloned repo](https://github.com/atanashristov/simplebank).

DB diagram at [dbdiagram](https://dbdiagram.io/d/6462ae5fdca9fb07c421e3e6).

Published DB docs at [dbdocs](https://dbdocs.io/atanashristov/simplebank).

Resources:

- [GirHub repo](https://github.com/techschool/simplebank)
- [DB Diagram web site](https://dbdiagram.io/home)
- [A tour of Go](https://go.dev/tour/welcome/1)

### Lecture 2: Install and use Docker and Postgres

Pull latest postgres image: `docker pull postgres:15.3-alpine`

Check the image is downloaded: `docker images`

Start container instance from the image: `docker run --name {container_name} -e {environment_variable} -p {host_ports:container_ports} -d {image:tag}`

Example: `docker run --name simplebank-postgres15 -p 5433:5432 -e POSTGRES_PASSWORD=post6res -d postgres:15.3-alpine`

Connect with pgsql to the database: `docker exec -it {container_name_or_id} {command} [args]`

Example: `docker exec -it simplebank-postgres15 psql -U postgres`

By default when connecting from local host postgres is set up on this image to not require password.

Stop container with: `docker stop {container_name}`

List running containers: `docker ps`

List all containers: `docker ps -a`

Start container: `docker start {container_name}`

Resources:

- [Postgres docker image](https://hub.docker.com/_/postgres)

### Lecture 3: Go migrations

Use the `migrate` library. They are multiple ways to [install migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

One possible way to install `migrate` is using the Go toolchain:

```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Verify installation: `migrate -help`

Commands:

- `create`: create new migration
- `goto V`: Migrate to version V
- `up [N]`: Apply all or N up migrations
- `down [N]`: Apply all or N down migrations

To apply the last 2 migrations: `migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2`

Create first migration from `simplebank` folder: `migrate create -ext sql -dir db/migration -seq init_schema`

Connect to the container shell: `docker exec -it simplebank-postgres15 /bin/sh`

Create new database: `createdb --username=postgres --owner=postgres simple_bank`

Connect to the new DB: `psql simple_bank -U postgres`

Delete the DB: `dropdb simple_bank -U postgres`

Another way to create database on the container is to directly run:

```sh
docker exec -it simplebank-postgres15 createdb --username=postgres --owner=postgres simple_bank
```

Then access the DB shell:

```sh
docker exec -it simplebank-postgres15 psql -U postgres simple_bank
```

**Makefile** - create `Makefile` in the root of the project, so you can run: `make dropdb`, `make createdb`, etc.

Note: On **Windows** install the `make` command with `choco install make`

SEE **Makefile** FOR ALL COMMANDS.

References:

- [Golang migrate library](https://github.com/golang-migrate/migrate)

### Lecture 4: Generate GRUD

Compare:

- [GORM](https://gorm.io/): complex and slow
- [sqlx](https://github.com/jmoiron/sqlx): faster and easy to use but not safe (catch typos at runtime)
- [sqlc](https://sqlc.dev/): fast, easy to use and safe. Generates go code from sql queries

Using `sqlc` in this course.

Install `sqlc` with:

```sh
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

Initialize _sqlc.yaml_ config file: `sqlc init`

Change the _sqlc.yaml_ like this:

```yaml
version: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: false # we make it `true` for optimization later
    emit_interface: false
    emit_exact_table_names: false # exact is accounts table -> "Accounts" struct, and we want "Account" struct
```

To generate the code run: `sqlc generate`.

**Windows** notes:

On Windows we have to use docker to generate the files, because on native Windows wqe get error:

Pull docker image: `docker pull kjconroy/sqlc`

Run sqlc using docker in Powershell: `docker run --rm -v ("$(pwd):/src" -replace '\\', '\\').ToLower() -w /src kjconroy/sqlc generate`

Created a powershell script: `.\sqlcw.ps1` for this.

You have to initialize a go module with: `go mod init github.com/atanashristov/simplebank`
, followed by `go mod tidy` to install any dependencies.

References:

- [sqlc settings](https://docs.sqlc.dev/en/latest/reference/config.html)
- [sqlc running with Docker](https://docs.sqlc.dev/en/stable/overview/install.html)

### Lecture 5: Unit tests for database CRUD

We need to install Postgres driver: `go get github.com/lib/pq`.

It adds it to the `go.mod` file:

```sh
require github.com/lib/pq v1.10.9 // indirect
```

This reference has to be imported in the go source files like this:

```go
import (
  _ "github.com/lib/pq"
)
```

In the tests we do not call directly any code from `lib/pq` but we have to include it, because we specify "postgres" as a `dbDriver` param to `sql.Open()`.

Then we run `go mod tidy` and the above reference in `go.mod` removed the "indirect".

To evaluate the test results we use the `testify` package.

Install `testify` with: `go get github.com/stretchr/testify`

Skip this: ... Then we run `go mod tidy` and the above reference in `go.mod` removed the "indirect".

Then import the `require` only in the `account_test.go`:

```go
import (
 "context"
 "testing"

  "github.com/stretchr/testify/require"
)
```

See:

- simplebank\db\sqlc\main_test.go
- simplebank\db\sqlc\account_test.go

References:

- [pq - A pure Go postgres driver for Go's database/sql package](https://github.com/lib/pq)
- [testify](https://github.com/stretchr/testify)

## Section 2: RESTful HTTP JSON Api (Gin + JWT + PASETO)

### Lecture 11: Implementing RESTful HTTP API in Go with Gin

Install Gin with: `go get -u github.com/gin-gonic/gin`.

Verify is installed in `simplebank\go.mod`.

Created API server in `simplebank\api\server.go`.

See:

[Gin GitHub page](https://github.com/gin-gonic/gin)

[Gin quick start](https://gin-gonic.com/docs/quickstart/)

[Gin model validation](https://gin-gonic.com/docs/examples/binding-and-validation/)

[Go validator, oneOf](https://pkg.go.dev/github.com/go-playground/validator#hdr-One_Of)
