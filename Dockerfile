# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

#ENTRYPOINT ["/app"]

# Command to run the executable
FROM alpine

WORKDIR /app

COPY --from=builder /app/main/ /app/

CMD ["./main"]