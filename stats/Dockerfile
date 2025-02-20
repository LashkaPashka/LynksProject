FROM golang:1.24.0

WORKDIR /app

COPY ./stats .

RUN go mod download

EXPOSE 8083

CMD ["go", "run", "cmd/main.go"]