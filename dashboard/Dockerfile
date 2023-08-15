# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy only the necessary Go mod files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o NeuraLaunch

# Stage 2: Create a minimal container
FROM alpine:latest

WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app/NeuraLaunch .

COPY .env .

# Expose the application port
EXPOSE 8080

# Define the default command
CMD ["./NeuraLaunch"]