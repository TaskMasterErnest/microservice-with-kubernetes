FROM golang:1.19 AS builder
ADD ./main.go main.go
ADD ./service service
# Copy the go.mod and go.sum files
COPY ./go.mod .
COPY ./go.sum .

# Fetch dependencies
RUN export GO111MODULE=on && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
