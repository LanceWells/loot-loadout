FROM golang:1.18.2 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY api api
COPY backend/characterimage backend/characterimage

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /app/app ./backend/characterimage/cmd/

EXPOSE 9999 40000

CMD ["dlv", "--listen=:40000","--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/app", "--continue"]
