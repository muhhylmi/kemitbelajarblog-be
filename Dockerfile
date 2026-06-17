FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# Run stage
FROM alpine:latest  

WORKDIR /app

# Add timezone data and CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Copy migrations to the path expected by runtime.Caller in main.go.
# In the builder, main.go is at /app/cmd/server/main.go.
# runtime.Caller resolves projectRoot to / (3 levels up from /app/cmd/server).
# So it expects migrations at /backend/migrations.
COPY --from=builder /app/migrations /backend/migrations

# Expose port
EXPOSE 3001

# Command to run the executable
CMD ["./server"]
