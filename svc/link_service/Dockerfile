FROM golang:1.20-alpine AS builder

WORKDIR $GOROOT/src/github.com/the-gigi/delinkcious/svc/link_service/service
# Copy files
COPY ./main.go main.go
COPY ./service service
COPY go.mod go.mod
COPY go.sum go.sum

# Fetch dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /link_service /app/link_service
EXPOSE 7070
ENTRYPOINT ["/app/link_service"]