# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Stage 2: Minimal image
FROM gcr.io/distroless/static
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

