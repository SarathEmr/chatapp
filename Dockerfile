# Start from the official Golang image
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o chatappserver .

# Use a minimal image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/chatappserver .
RUN chmod +x chatappserver

# Expose port 3000
EXPOSE 3000

# Command to run the server
CMD ["./chatappserver"]