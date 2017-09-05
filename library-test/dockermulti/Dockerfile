FROM golang:1.9 as builder
WORKDIR /go/src/go-learning/dockermulti/ares
ADD common /go/src/go-learning/dockermulti/common
ADD ares /go/src/go-learning/dockermulti/ares
RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o ares .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --update bash && rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=builder /go/src/go-learning/dockermulti/ares/ares .
ENTRYPOINT ["./ares"]
