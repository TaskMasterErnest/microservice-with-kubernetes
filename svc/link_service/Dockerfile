FROM golang:alpine AS builder

WORKDIR $GOROOT/src/github.com/TaskMasterErnest/microservice-with-kubernetes/svc/link_service
COPY ./main.go main.go
COPY ./service service

# Fetch dependencies
RUN go mod init link_service && \
    go get && \
    CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /link_service /app/link_service
EXPOSE 7070
ENTRYPOINT ["/app/link_service"]
