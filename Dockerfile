FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o microblog main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/microblog .

EXPOSE 8080

CMD ["./microblog"]