# ctf01d-training-platform

Service used go && psql

[http://localhost:4102/](http://localhost:4102/)

## Install requriments

```shell
$ sudo apt install golang
$ go mod download
```

## Build server

```shell
$ cd src
$ go build main.go
```

## Run local dev server

```shell
$ cd src
$ go run main.go
```

## Run DB for local devlopment

### psql -- actual

connect to db configure in `src/config/config.yaml`

```yaml
...
db:
  driver: postgres
  data_source: postgres://postgres@localhost:5432/ctf01d
...
```

run local container

```shell
$ docker run -d --name ctf01d-postgres -e POSTGRES_DB=ctf01d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
```

attach to container

```shell
$ docker exec -it ctf01d-postgres psql -U postgres
```
