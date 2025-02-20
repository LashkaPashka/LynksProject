FROM golang:1.24.0

WORKDIR /app

COPY ./memcache .

RUN go mod download

EXPOSE 8084

CMD ["go", "run", "cmd/main.go"]