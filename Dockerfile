FROM golang:alpine AS builder

RUN apk add --no-cache git && CGO_ENABLED=0 GOOS=linux go get github.com/alash3al/httpsify

FROM scratch

COPY --from=builder /go/bin/httpsify /go/bin/httpsify

ENTRYPOINT ["/go/bin/httpsify"]
