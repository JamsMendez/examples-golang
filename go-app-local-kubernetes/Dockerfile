FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o go-app

FROM ubuntu:24.10 AS production

WORKDIR /app

COPY --from=builder /app/go-app .

EXPOSE 3000

CMD ["./go-app"]
