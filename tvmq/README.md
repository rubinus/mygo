# tvmq
###this is tvmining.com queue for message

fnsq文件夹是用来消费nsq队列用的，比如新用户加入，聊天消息

kafka是用来消费后台推送事件用的，如果新加，只需要复制一份代码，改相应该的逻辑，然后加到kafka/run.go中

开发测试时可以更改 Dockerfile文件 的pro/dev/qa，对应config下的json配置


mac机器执行下面的，生成image
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tvmq

cd tvmq

docker build -t rubinus/tvmq .

docker tag rubinus/tvmq dkhub.tvmining.com:5000/tvmq:1.0.0

docker push  dkhub.tvmining.com:5000/tvmq:1.0.0


docker service update --image dkhub.tvmining.com:5000/tvmq:1.0.0 serv-cluster_socket

回滚错误的更新
docker service rollback serv-cluster_socket