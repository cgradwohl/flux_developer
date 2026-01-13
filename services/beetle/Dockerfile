# ---------- build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for go modules sometimes)
RUN apk add --no-cache git

# Copy go mod files first for better layer caching
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy source
COPY src/ .

# Build statically-linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o beetle .

# ---------- runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/beetle /beetle

# Run as non-root for safety
USER nonroot:nonroot

ENTRYPOINT ["/beetle"]

