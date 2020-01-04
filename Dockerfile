FROM golang:1.13 AS builder
WORKDIR /marvin
COPY go.mod /marvin
COPY go.sum /marvin
COPY src/ /marvin
RUN CGO_ENABLED=0 GOOS=linux go build -a -o marvin .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /marvin .
CMD ["./marvin"]