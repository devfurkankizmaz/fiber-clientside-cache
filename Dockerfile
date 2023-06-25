# Builder stage
FROM golang:1.19 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the project files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the port
EXPOSE 3000

# Run the application
CMD ["./app"]
