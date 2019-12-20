FROM golang:1.13-alpine as builder

ARG port
ENV PORT=$port

WORKDIR /dist
COPY . .

ARG src
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $src

FROM busybox:latest

WORKDIR /dist
COPY --from=builder /dist .
CMD ["./server"]
