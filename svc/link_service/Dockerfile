FROM golang:1.20-alpine AS builder

WORKDIR $GOROOT/src/github.com/the-gigi/delinkcious/svc/link_service
# Copy files
COPY ./main.go main.go
COPY ./service service

# Update
RUN apk update && apk upgrade && \
    go mod init link_service

# Fetch dependencies
RUN go get && \
  # Build image as a truly static Go binary
    CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /link_service /app/link_service
EXPOSE 8080
ENTRYPOINT ["/app/link_service"]