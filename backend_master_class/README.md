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

Note: On **Windows** instll the `make` command with `choco install make`

SEE **Makefile** FOR ALL COMMANDS.

References:

- [Golang migrate library](https://github.com/golang-migrate/migrate)
