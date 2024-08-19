# Stage 1: Build the application
FROM golang:1.21.8-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Cache go mod dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy only the necessary directories
COPY cmd/ ./cmd/
COPY config/ ./config/
COPY data/ ./data/
COPY pkg/ ./pkg/
COPY internal/ ./internal/

# Build the Go app
RUN go build -o api ./cmd/app

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy data file
COPY --from=builder /app/data/meteorite-landings.json /app/data/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

# Expose port 8083 to the outside world
EXPOSE 8083

# Command to run the executable
CMD ["./api"]
