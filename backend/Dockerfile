# Start with a base Golang builder image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
