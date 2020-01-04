# Start from golang base image
FROM golang:alpine as builder

# Set the current working directory inside the container
WORKDIR /dist

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

# Build the Go app
ARG path
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $path

# Start a new stage from busybox
FROM busybox:latest

WORKDIR /dist

# Copy the pre-built binaries and the env file from the previous stage
COPY --from=builder /dist/server .

# Run the executable
CMD ["./server"]
