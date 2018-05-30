FROM golang:latest
LABEL maintainer="rubinus.chu@gmail.com"
#设置工作目录
WORKDIR $GOPATH/src/mygo
ADD . $GOPATH/src/mygo
#下载第三方包
RUN go get github.com/json-iterator/go && go get github.com/Shopify/sarama && \
go get github.com/bsm/sarama-cluster && go get github.com/garyburd/redigo/redis && \
go get github.com/mailru/easyjson && go get github.com/mailru/easyjson/jlexer && \
go get github.com/mailru/easyjson/jwriter && go get gopkg.in/mgo.v2 && \
go get gopkg.in/mgo.v2/bson
#go构建可执行文件
RUN go build .
#最终运行docker的命令
ENTRYPOINT  ["./mygo"]