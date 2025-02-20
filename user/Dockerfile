FROM golang:1.24.0

WORKDIR /app

COPY ./user .

RUN go mod download

EXPOSE 8082

CMD ["go", "run", "cmd/main.go"]