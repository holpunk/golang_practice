# Use an official Go image as the base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Download all dependencies
RUN go mod init api
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Use a lightweight Alpine Linux base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable when the container starts
CMD ["./main"]
