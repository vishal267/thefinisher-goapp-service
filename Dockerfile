# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Stage 2: Minimal image
FROM gcr.io/distroless/static
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

