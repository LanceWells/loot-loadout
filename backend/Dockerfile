FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./service_public ./service_public

RUN go build -o /loot-loadout-backend ./cmd/chatsocket/main.go
EXPOSE 8080

CMD ["/loot-loadout-backend"]
