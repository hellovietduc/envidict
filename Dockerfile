# Build stage
FROM golang:1.14-alpine AS build

WORKDIR /go/src/github.com/vietduc01100001/envidict
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/envidict ./cmd/envidict

# Final image
FROM alpine

ENTRYPOINT ["/go/bin/envidict"]
ENV GIN_MODE=release

COPY ./etc/en-vi-dict.txt /etc/envidict/en-vi-dict.txt
COPY --from=build /go/bin/envidict /go/bin/envidict
COPY ./static /etc/envidict/static
