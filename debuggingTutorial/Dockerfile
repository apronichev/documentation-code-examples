FROM golang:1.14.3-alpine3.11 AS build-env

ENV CGO_ENABLED 0

# Allow Go to retreive the dependencies for the build step
RUN apk add --no-cache git

WORKDIR /debuggingTutorial/
ADD . /debuggingTutorial/

RUN go build -o /debuggingTutorial/srv .

# Get Delve from a GOPATH not from a Go Modules project
WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

# final stage
FROM alpine:3.11

WORKDIR /
COPY --from=build-env /debuggingTutorial/srv /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/srv"]