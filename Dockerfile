# Start from golang base image
FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

COPY . .

ADD . /go/src/app

RUN go get -v