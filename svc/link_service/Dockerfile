FROM golang:1.19 AS builder
ADD ./main.go main.go
ADD ./service service
# Fetch dependencies
RUN export GO111MODULE=on && \
    go get -dv -m

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /link_service /app/link_service
EXPOSE 7070
ENTRYPOINT ["/app/link_service"]
