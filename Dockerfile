# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app/
RUN mkdir ./logs

# Copy the go.mod and go.sum file into your project directory /app
COPY go.mod go.sum ./ 

# The go mod download command downloads the named modules into the module cache.
RUN go mod download 

# Copy everything in the current directory into a Docker image
COPY . /app/

WORKDIR /app/cmd/app

# Compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-coding-challenges .

# Run
CMD ["/docker-coding-challenges"]