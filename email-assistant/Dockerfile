FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -buildvcs=false -o main ./src/main.go

WORKDIR /dist

RUN cp /build/main .

FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

COPY --from=builder /dist/main /main

EXPOSE 8000

ENTRYPOINT ["/main"]
