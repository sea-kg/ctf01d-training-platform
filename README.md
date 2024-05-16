# ctf01d-training-platform

Service used go && psql

[http://localhost:4102/](http://localhost:4102/)

## Install requriments

```shell
$ go mod download
```

## Build server

```shell
$ go build cmd/ctf01d/main.go
```

## Run local dev server

```shell
$ go run cmd/ctf01d/main.go
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

## Hints

fuzz api

```shell
docker run --net=host --volume /home/user/ctf01d-training-platform/api:/api/ ghcr.io/matusf/openapi-fuzzer run -s '/api/s
wagger.yaml'  --url http://localhost:4102


docker run --net=host  --volume /home/user/ctf01d-training-platform/api:/api/ kisspeter/apifuzzer --src_file '/api/swagger.yaml'  --url http://localhost:4102 -r /api/
```
