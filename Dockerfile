FROM golang:1.21-alpine AS build_base

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-app

COPY . .

# Build the Go app
RUN go build -o ./bin/transaction cmd/api/main.go

# Start fresh from a smaller image
FROM alpine:3.19
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-app/bin/transaction /app/transaction

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/transaction"]