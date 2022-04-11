FROM golang:1.18

ADD . /home

WORKDIR /home

# RUN \
#        apk add --no-cache bash git openssh && \
#        go get -u github.com/minio/minio-go


CMD ["go","run","main.go"]