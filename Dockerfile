FROM golang:1.22-bookworm as builder
LABEL "maintainer"="Evgenii Sopov <mrseakg@gmail.com>"
LABEL "repository"="https://github.com/sea-kg/ctf01d"

WORKDIR /go/src/app

COPY ./src/go.mod ./

RUN go mod tidy

COPY ./src/ ./

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 4202

CMD ["./server"]

