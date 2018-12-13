FROM alpine:latest
LABEL maintainer="zhuhonglei@tvmining.com"

#访问https
RUN apk add --no-cache ca-certificates

#设置工作目录
WORKDIR $GOPATH/src/serv/
ADD tvmq $GOPATH/src/serv/
ADD static $GOPATH/src/serv/static
ADD config $GOPATH/src/serv/config
ADD templates $GOPATH/src/serv/templates

RUN ["chmod", "+x", "tvmq"]

EXPOSE 8081

ENTRYPOINT ["./tvmq","--env","qa"]