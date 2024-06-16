FROM golang:1.19-alpine AS builder

WORKDIR /app

# Install Glide as a dependency manager (optional, replace with preferred method)
RUN apk add --no-cache curl && \
    curl -sSL https://github.com/Masterminds/glide/releases/download/v0.13.7/glide-v0.13.7-x86_64-linux-musl.tar.gz | tar -xzf -

# Copy your application code
COPY go.mod glide.yaml ./

# Install dependencies using Glide (replace with preferred method)
RUN glide install

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
CMD ["/app/main"]