# Build stage
FROM golang:1.25.1-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations for Cloud Run
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main .

# Final stage - use distroless for smaller size and security
FROM gcr.io/distroless/static-debian12:nonroot

# Copy ca-certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary from builder stage
COPY --from=builder /app/main /app/main

# Use non-root user (already set in distroless/static:nonroot)
# USER 65534

# Set environment variables for Cloud Run
ENV GIN_MODE=release
ENV PORT=8080

# Cloud Run will set the PORT environment variable
# The application should listen on 0.0.0.0:$PORT

# Expose port (Cloud Run ignores this but it's good for documentation)
EXPOSE 8080

# Health check endpoint (optional but recommended for Cloud Run)
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

# Run the application
CMD ["/app/main"]