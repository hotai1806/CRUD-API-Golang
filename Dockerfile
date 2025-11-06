FROM golang:1.25-alpine

# Install git and air
RUN apk add --no-cache git && go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files including .air.toml
COPY . .
COPY .air.toml .

# Expose app port
EXPOSE 8080

# Run Air with config
CMD ["air", "-c", ".air.toml"]
