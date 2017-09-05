FROM golang:1.9

# Copy the local package files to the container’s workspace.
ADD common /go/src/hue-ec-swat-golang/common
ADD ares /go/src/hue-ec-swat-golang/ares
ADD harmonia /go/src/hue-ec-swat-golang/harmonia

# Install dependencies
RUN go get gopkg.in/yaml.v2
RUN go get github.com/streadway/quantile

# Install project binary gloablly within container
#RUN go install hue-ec-swat-golang/ares
RUN go build hue-ec-swat-golang/ares

# Initial commands to run when image run
#CMD cd /go/src/hue-ec-swat-golang/ares

# Set binary as entrypoint
#ENTRYPOINT /go/bin/ares
ENTRYPOINT  ["/go/src/hue-ec-swat-golang/ares/ares"]

# Expose default port (3000)
EXPOSE 3000