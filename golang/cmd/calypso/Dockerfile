FROM golang:1.13 as builder

ENV PORT=3001

WORKDIR /dist
COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./golang/cmd/calypso/

FROM scratch

WORKDIR /dist
COPY --from=builder /dist .
CMD ["./server"]
