FROM golang:1.9

# Copy the local package files to the container’s workspace.
ADD hello /go/src/go-learning/hello

# Install project binary gloablly within container
#RUN go install hue-ec-swat-golang/ares
RUN go build go-learning/hello

# Initial commands to run when image run
#CMD cd /go/src/hue-ec-swat-golang/ares

# Set binary as entrypoint
#ENTRYPOINT /go/bin/ares
ENTRYPOINT  ["/go/src/go-learning/hello/hello"]
