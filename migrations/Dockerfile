# Base image
FROM alpine:latest AS goosebase

# grab curl
RUN apk add --no-cache curl

# Download and install Goose 
RUN curl -L -o /bin/goose https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64

# Final image
FROM alpine:latest

# Copy the Goose binary from the base stage
COPY --from=goosebase /bin/goose /bin/goose
RUN chmod +x /bin/goose

WORKDIR /app

# copy migration files to the container
COPY ./schemas /app/schemas