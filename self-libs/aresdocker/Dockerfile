FROM golang:ares as builder
WORKDIR /go/src/hue-ec-swat-golang/ares/docker
RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o docker .

FROM alpine:go
#RUN apk add --update bash && rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=builder /go/src/hue-ec-swat-golang/ares/docker/docker .
ENTRYPOINT ["./docker"]