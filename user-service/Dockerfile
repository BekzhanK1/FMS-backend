# Stage 1: Build the application
FROM golang:1.22.3-bullseye AS builder

# Install required tools for building
RUN apt-get update && apt-get install -y build-essential

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files for dependency resolution and building
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod vendor

# Set GOFLAGS globally for consistent usage
ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

# Copy the entire project directory, including the vendor directory
COPY . .

# Build the application binary
RUN make build

# Stage 2: Create a smaller image for running the application
FROM alpine:3.18

# Install runtime dependencies
RUN apk add --no-cache ca-certificates bash

# Set the working directory
WORKDIR /app

# Copy the built binary and necessary scripts from the builder stage
COPY --from=builder /app/bin/user-service /app/user-service
COPY --from=builder /app/wait-for-it.sh /app/wait-for-it.sh
COPY --from=builder /app/migration.sh /app/migration.sh

# Ensure scripts have execute permissions (correct paths inside /app)
RUN chmod +x /app/wait-for-it.sh /app/migration.sh

# Expose the application's port
EXPOSE 5002

# Command to run migrations and start the application
ENTRYPOINT ["/bin/bash", "-c", "/app/wait-for-it.sh postgres:5432 -- /app/migration.sh && ./user-service"]
