# Build Stage
FROM golang:1.20.3-alpine3.17 AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY . ./

# Do the migration
RUN go run migrate/migrate.go

# Build the binary.
RUN go build -v -o main .

# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 9090

# Run the service
CMD ["./main"]