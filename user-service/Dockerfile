# Use the official Go image for building the application
FROM golang:1.22.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files for dependency resolution and building
COPY go.mod go.sum ./

# Enable vendoring and download dependencies into the vendor directory
RUN go mod tidy && go mod vendor

# Set GOFLAGS globally for consistent usage
ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

# Copy the entire project directory, including the vendor directory
COPY . .

# Use the Makefile to build the application
RUN make build

# Create a smaller image for running the application
FROM alpine:3.18

# Install required certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/user-service /app/user-service

# Expose the application's port (modify as needed)
EXPOSE 5002

# Command to run the application
ENTRYPOINT ["./user-service"]
