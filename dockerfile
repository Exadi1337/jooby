FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download 
COPY . .
RUN go build -o jooby ./main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/jooby .
CMD ["./jooby"]
