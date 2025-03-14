FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main

FROM golang:1.23
WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]
