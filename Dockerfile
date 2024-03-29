FROM golang:latest
LABEL maintainer="rubinus.chu@gmail.com"
#设置工作目录

ADD github.com/ $GOPATH/src/github.com/
ADD golang.org/ $GOPATH/src/golang.org/
ADD gopkg.in/ $GOPATH/src/gopkg.in/
ADD mygo/ $GOPATH/src/mygo/

WORKDIR $GOPATH/src/mygo

#RUN go get -d -v ./...

#下载第三方包
#RUN go get -v gopkg.in/mgo.v2 && go get -v gopkg.in/mgo.v2/bson
#RUN go get github.com/json-iterator/go && go get github.com/garyburd/redigo/redis && \
#go get github.com/mailru/easyjson
#RUN go get github.com/json-iterator/go && go get github.com/Shopify/sarama && \
#go get github.com/bsm/sarama-cluster && go get github.com/garyburd/redigo/redis && \
#go get github.com/mailru/easyjson && go get github.com/mailru/easyjson/jlexer && \
#go get github.com/mailru/easyjson/jwriter && go get gopkg.in/mgo.v2 && \
#go get gopkg.in/mgo.v2/bson && go get github.com/nsqio/go-nsq
#go构建可执行文件
#RUN proxychains4 go get golang.org/x/text/
#RUN go build .
#最终运行docker的命令

EXPOSE 8080
#ENTRYPOINT ["./mygo"]
CMD ["go","run","main.go"]