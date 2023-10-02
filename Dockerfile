# Create build stage based on buster image
FROM golang:latest AS builder
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY . .
# Install any required modules
RUN go mod download
# Copy over Go source code
COPY . .
# Run the Go build and output binary under hello_go_http
RUN go build -o /hello_go_http

# Run the app binary when we run the container
ENTRYPOINT ["/hello_go_http"]