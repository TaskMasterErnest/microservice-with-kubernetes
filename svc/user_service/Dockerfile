FROM golang:1.17 AS builder

WORKDIR $GOROOT/src/github.com/TaskMasterErnest/microservice-with-kubernetes/svc/user_service
# Copy files
COPY ./main.go main.go
COPY ./service service
# Copy files
COPY ./go.mod go.mod

# Fetch dependencies
RUN go get && \
    CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /user_service /app/user_service
EXPOSE 7070
ENTRYPOINT ["/app/user_service"]
