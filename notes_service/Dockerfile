# Builder
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy the Go mod
COPY go.mod go.sum ./

# Download deps
RUN go mod download

# Copy app code
COPY . .

# Build binary
RUN go build -o main .

# Runner
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the server port
EXPOSE 8080

# Run the compiled binary
CMD ["./main"]