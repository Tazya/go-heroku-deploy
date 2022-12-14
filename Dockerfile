FROM golang:1.19-alpine as builder
WORKDIR /build

ADD go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go build -v -mod=vendor -o app main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /build/app /usr/local/bin/

CMD /usr/local/bin/app --port $PORT
