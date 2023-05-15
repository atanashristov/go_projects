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

Resources:

- [Postgres docker image](https://hub.docker.com/_/postgres)
