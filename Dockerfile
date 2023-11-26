FROM golang:alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY ./ ./


RUN go build ./cmd/avito
EXPOSE 8090
# Запустим приложение
CMD ["./avito"]