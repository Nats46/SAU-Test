FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./main.go

FROM alpine:latest

COPY --from=builder /app/myapp /usr/local/bin/myapp

CMD ["myapp"]

EXPOSE 8080
