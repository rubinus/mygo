FROM alpine:latest
LABEL maintainer="zhuhonglei@tvmining.com"
#设置工作目录
WORKDIR $GOPATH/src/serv/
ADD ./static/headimgurl/genuser $GOPATH/src/serv/
ADD static $GOPATH/src/serv/static
ADD config $GOPATH/src/serv/config
RUN ["chmod", "+x", "genuser"]
ENTRYPOINT ["./genuser","--env","pro"]