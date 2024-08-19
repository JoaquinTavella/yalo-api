# Stage 1: Build the Go application
FROM golang:1.22.6-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Build the Go app, using static linking to create a binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Set the working directory inside the final image
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/app .

EXPOSE 8080

# Run the Go app
CMD ["./app"]
