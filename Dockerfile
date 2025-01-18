# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Set up the working directory inside the container
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN go build -o arrivals-service .

# Stage 2: Create a minimal image with the compiled binary
FROM alpine:latest

# Set up the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/arrivals-service .

# Expose the port your Go service will listen on
EXPOSE 8080

# Run the Go binary
CMD ["./arrivals-service"]
