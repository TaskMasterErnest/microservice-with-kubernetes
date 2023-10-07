FROM golang:1.17-alpine AS builder

WORKDIR $GOROOT/src/github.com/the-gigi/delinkcious/svc/social_graph_service
# Copy files
COPY ./main.go main.go
COPY ./service service
COPY go.mod go.mod
COPY go.sum go.sum

# Fetch dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
