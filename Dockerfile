# Stage 1: Build stage
FROM golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-klaviyo-to-bigquery .

# Stage 2: Final stage
FROM alpine

# Echo the version of the image
RUN echo "go-klaviyo-to-bigquery v0.1.0"
# Set the working directory
WORKDIR /app


# Copy the binary from the build stage
COPY --from=build /app/go-klaviyo-to-bigquery .
COPY --from=build /app/config.yaml .
COPY --from=build /app/.local .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/go-klaviyo-to-bigquery", "--config", "/app/config.yaml"]