FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install required dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hms ./cmd/main

# Use a small alpine image
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/hms .
COPY --from=builder /app/migrations ./migrations

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./hms"]
