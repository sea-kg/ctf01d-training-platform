FROM golang:1.15-buster
LABEL "maintainer"="Evgenii Sopov <mrseakg@gmail.com>"
LABEL "repository"="https://github.com/sea-kg/ctf01d"

WORKDIR /go/src/app

COPY ./src/ /go/src/app

# Better use a localfolders
RUN go install github.com/go-sql-driver/mysql@latest
RUN go install github.com/jmoiron/sqlx@latest
RUN go install github.com/gorilla/mux@latest

EXPOSE 4202

CMD exec go run server.go

# CMD ["go","run","server.go"]


