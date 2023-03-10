############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --update 

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/visiongateway/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v ./...

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main 

############################
# STEP 2 build a small image
############################
FROM alpine:3

WORKDIR /

# Copy our static executable.
COPY --from=builder /main /go/main

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

WORKDIR /go
COPY /config.json /go

# Run the Go Gin binary.
ENTRYPOINT ["/go/main"]
