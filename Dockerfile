# syntax=docker/dockerfile:1
# Build the application from source
FROM golang:1.19 AS builder
# Set destination for ADD
WORKDIR /ports
# Copy the source code
ADD . .
# Download Go modules
RUN go mod download
RUN go mod verify
# Run tests
RUN go test -v ./...
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/ports/main.go

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release
RUN apk --no-cache --update add ca-certificates
# Copy app binary from builder image
COPY --from=builder /bin/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
# Run
CMD ["app"]
