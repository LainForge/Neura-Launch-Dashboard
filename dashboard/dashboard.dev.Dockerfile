# Build the application
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy only the necessary Go mod files
COPY go.mod .
COPY go.sum .

# Download dependencies

RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose the application port
EXPOSE 8080

# Define the default command
CMD ["go", "run", "main.go"]