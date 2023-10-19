FROM golang:1.20-alpine AS builder

WORKDIR /app
RUN mkdir "/build"

# Copy the necessary files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose PORT
EXPOSE 5000

# Run the application 
ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"


