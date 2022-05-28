FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./*.go .

RUN go build -o /roomsocket ./cmd/main.go
EXPOSE 8080

CMD ["/roomsocket"]
