FROM golang:1.13 AS builder
WORKDIR /marvin
COPY go.mod .
COPY go.sum .
RUN mkdir src
COPY src/ ./src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o marvin ./src

# TODO: trust .dockerignore and copy less explicitly
FROM node:13.5-alpine AS frontend-builder
WORKDIR /frontend
RUN mkdir src
RUN mkdir public
COPY assets/src/ ./src/
COPY assets/public/ ./public/
COPY assets/jsconfig.json .
COPY assets/package.json .
COPY assets/yarn.lock .
RUN npm install
RUN npm run build

# TODO: environment variable to specify build folder
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir assets/
COPY --from=builder /marvin .
COPY --from=frontend-builder /frontend/build/ assets/build
CMD ["./marvin"]