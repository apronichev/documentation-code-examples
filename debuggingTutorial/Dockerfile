# Use a more recent version of Go
FROM golang:1.21-alpine AS build-env

ENV CGO_ENABLED=0

# Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache git

WORKDIR /debuggingTutorial/
COPY . .

# Build the application
RUN go build -o /debuggingTutorial/app .

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Final stage
FROM alpine:3.18

WORKDIR /
COPY --from=build-env /debuggingTutorial/app /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/app"]