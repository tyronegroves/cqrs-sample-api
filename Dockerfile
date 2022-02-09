FROM golang:1.18beta2 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go .
ADD config ./config
ADD server ./server

RUN CGO_ENABLED=0 go build -o api

FROM scratch

COPY --from=builder /app/api /api
COPY --from=builder /app/config /config

ENTRYPOINT ["/api"]