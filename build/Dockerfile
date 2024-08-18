FROM golang:1.23-bookworm as builder

WORKDIR /ctf01d.ru

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod tidy

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

FROM alpine:latest as prod
LABEL "maintainer"="Evgenii Sopov <mrseakg@gmail.com>"
LABEL "repository"="https://github.com/sea-kg/ctf01d"
RUN apk --no-cache add ca-certificates

WORKDIR /ctf01d.ru

COPY --from=builder /ctf01d.ru/server .
COPY --from=builder /ctf01d.ru/configs/ ./configs/
COPY --from=builder /ctf01d.ru/html/ ./html/

EXPOSE 4102

CMD ["./server"]
