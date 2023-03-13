# Use the official Golang image as builder
FROM golang:1.20.1-alpine3.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application binary inside the container
RUN go build -o main .

# Use a lightweight Alpine Linux image as the final base image
FROM alpine:3.17

# Set the working directory in the final image
WORKDIR /app

# Copy the binary from the builder image to the final image
COPY --from=builder /app/main .

# Set the command to run the binary
CMD ["./main"]
