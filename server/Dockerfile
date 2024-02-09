# Stage 1: Build the Go binary
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set necessary environment variables
ENV PORT=8080

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port
EXPOSE $PORT

# Command to run the executable
CMD ["./myapp"]