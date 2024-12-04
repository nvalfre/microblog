FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microblog main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/microblog .

RUN chmod +x /microblog

EXPOSE 8080

CMD ["./microblog"]