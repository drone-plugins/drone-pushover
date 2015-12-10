# Docker image for the Drone Pushover plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-pushover
#     make deps build
#     docker build --rm=true -t plugins/drone-pushover .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-pushover /bin/
ENTRYPOINT ["/bin/drone-pushover"]
