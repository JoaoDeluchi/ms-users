
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy your application code
COPY go.mod ./

# Download and install dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api/main.go

# Stage 2: Production image (slim Alpine with application binary)
FROM alpine:latest

WORKDIR /app

# Copy application binary from builder stage
COPY --from=builder /app/main /app/main

# Expose port for the Gin server
EXPOSE 8080

# Command to run the application
CMD ["/cmd/api/main"]