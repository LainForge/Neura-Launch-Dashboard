# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
RUN mkdir "/build"


# Copy only the necessary Go mod files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go mod download

# Copy the rest of the source code
COPY . .
COPY ./.env .

EXPOSE 8080

ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"
