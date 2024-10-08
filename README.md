# Backend for ctf01d-training-platform

This service uses Go and PostgreSQL.

[Roadmap](docs/ROADMAP.md) & [Concept](docs/CONCEPT.md)

### Related Projects
- [Frontend for ctf01d-training-platform](https://github.com/AlexBazh/ctf01d-front/)

### Contributing
We welcome contributions! Please see our [Contributing Guidelines](docs/CONTRIBUTING.md) for more details.

## Project structure

```sh
.
├── api # openapi schema
│   └── openapi.yaml
├── build # docker files
│   ├── docker-compose.stage.yml
│   ├── docker-compose.yml
│   └── Dockerfile
├── cmd # main app
│   └── main.go
├── configs # app config and other config (linter)
│   ├── config.development.yml
│   ├── config.production.yml
│   ├── config.test.yml
│   ├── golangci
│   │   └── golangci.yml
│   └── spectral
│       └── spectral.yaml
├── docs # project documentation
│   └── *.md
├── html # current frontend
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler # backend handlers
│   │   └── *.go
│   ├── helper # shared code
│   │   └── helper.go
│   ├── logger
│   │   └── logger.go
│   ├── migrations # sql migrations
│   │   └── psql
│   │       ├── struct_updater.go
│   │       └── update*.go
│   ├── model # api json template
│   │   └── *.go
│   ├── repository # work with db, query ...
│   │   └── *.go
│   ├── httpserver # autogenerated code from oapi-codegen
│   │   ├── html_spa.go
│   │   └── httpserver.gen.go
│   └── view
│       └── session.go
├── pkg
│   └── avatar # uniq pic for user
│       └── avatar.go
├── scripts
│   └── create-migration.go # make new file for migration
└── test
    └── server_integration_test.go
```

## Development

### Install go on `Ubuntu 24.04 LTS`

```shell
$ snap install go --classic
```

also you will need docker

### Build and start from source code

```shell
$ make run-db
...
$ make build
$ ./main
```

And then open
http://localhost:4102/

Default admin credentials: admin/admin


### Install requirements

```shell
$ go mod download && go mod tidy
```

### Build server

```shell
$ go build cmd/main.go
```

### Run server

```shell
$ go run cmd/main.go
```

will be available on - [http://localhost:4102](http://localhost:4102)


### Generate code from openapi schema

```shell
oapi-codegen -generate models,chi -o internal/server/httpserver.gen.go --package routers api/openapi.yaml
```

## DataBase

### psql

connect to db configure in `src/configs/config.#{STAGE}.yaml`

```yaml
...
db:
  driver: postgres
  data_source: postgres://postgres@localhost:5432/ctf01d
...
```

### run db local container

```shell
$ docker run -d --name ctf01d-postgres -e POSTGRES_DB=ctf01d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
```

### attach to db container

```shell
$ docker exec -it ctf01d-postgres psql -U postgres
```

### create new migrate file

```shell
go run scripts/create-migration.go # Created new migration file: internal/migrations/psql/update0022_update0023.go
```

## Experimental

### fuzz api

```shell
docker run --net=host --volume /home/user/ctf01d-training-platform/api:/api/ ghcr.io/matusf/openapi-fuzzer run -s '/api/openapi.yaml' --url http://localhost:4102

docker run --net=host --volume /home/user/ctf01d-training-platform/api:/api/ kisspeter/apifuzzer --src_file '/api/openapi.yaml' --url http://localhost:4102 -r /api/
```


## Generate Go server boilerplate from OpenAPI 3 - oapi-codegen

install:

```shell
$ go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
$ export PATH="$PATH:$HOME/bin:$HOME/go/bin"
```


## Database local

```shell
$ make run-db
$ psql postgresql://postgres:postgres@localhost:4112/ctf01d_training_platform
ctf01d_training_platform=#
```
