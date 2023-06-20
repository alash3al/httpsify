FROM golang:1.18-alpine As builder

WORKDIR /httpsify/

RUN apk update && apk add git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /usr/bin/httpsify .

FROM alpine

WORKDIR /httpsify/

COPY --from=builder /usr/bin/httpsify /usr/bin/httpsify

CMD httpsify