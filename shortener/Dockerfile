FROM golang:1.24.0

WORKDIR /app

COPY ./shortener .

RUN go mod download

EXPOSE 8081

CMD ["go", "run", "cmd/main.go"]