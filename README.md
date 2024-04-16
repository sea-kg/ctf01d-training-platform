# ctf01d-training-platform

Service used go && mysql

[http://localhost:4102/](http://localhost:4102/)

## Install requriments

```shell
$ sudo apt install golang
$ go get github.com/go-sql-driver/mysql
$ go get github.com/jmoiron/sqlx
$ go get github.com/gorilla/mux
```

## Build server

```shell
$ cd src
$ go build server.go
```

## Run local dev server

```shell
$ cd src
$ SERVICE2_GO_MYSQL_HOST=service2_go_db SERVICE2_GO_MYSQL_DBNAME=service2_go SERVICE2_GO_MYSQL_USER=service2_go SERVICE2_GO_MYSQL_PASSWORD=service2_go go run server.go
```

## Run DB for local devlopment

```shell
$ docker run --name service2_go_db -e MYSQL_ROOT_PASSWORD=service2_go -e MYSQL_DATABASE=service2_go -e MYSQL_USER=service2_go -e MYSQL_PASSWORD=service2_go -p 3306:3306 -d mysql:latest
```
