FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal ./internal
COPY cmd ./cmd

RUN go build -o /app/tg-bot ./cmd/

EXPOSE 8080

CMD ["/app/tg-bot"]
