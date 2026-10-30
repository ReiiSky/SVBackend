FROM golang:1.26.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o SVBackend cmd/rest/*.go

# Runtime stage
FROM alpine:edge

# Set working directory
WORKDIR /app

# Copy built binary from build stage
COPY --from=build /app/SVBackend .

# Install runtime dependencies
RUN apk --no-cache add \
    ca-certificates \
    tzdata

# Run application
CMD ["/app/SVBackend"]
