FROM gliderlabs/alpine:latest

RUN apk --update add ca-certificates bash

# See: http://www.blang.io/posts/2015-04_golang-alpine-build-golang-binaries-for-alpine-linux/
# And: https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
# CGP RUN apk --update add ca-certificates gcc libc-dev
# https://github.com/jamiemccrindle/dockerception

# GOGC=off CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#ADD main /
COPY main /
COPY assets /assets/
COPY etc /etc/
COPY img /img/
COPY tpl /tpl/

ENV VIRTUAL_HOST "eutopia.fail,www.eutopia.fail"
ENV VIRTUAL_PORT 8080

EXPOSE 8080

CMD ["/main"]

# docker build -t eu:v1 .
# docker save eu:v1 | bzip2 | pv | ssh root@de.allancorfix.dk 'bunzip2 | docker load'