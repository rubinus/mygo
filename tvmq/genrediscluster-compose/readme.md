在当前目录下运行
docker build -t rubinus/redis-cluster .

docker-compose up -d

docker ps

登录任意一个容器执行
redis-cli --cluster create 10.20.80.132:7001 10.20.80.132:7002 10.20.80.132:7003 10.20.80.132:7004 10.20.80.132:7005 10.20.80.132:7006 --cluster-replicas 1

以下是redis5以下version建立集群时用的
docker run --rm -it --net host inem0o/redis-trib create --replicas 1 192.168.2.1:7001 192.168.2.1:7002 192.168.2.1:7003 192.168.2.1:7004 192.168.2.1:7005 192.168.2.1:7006

说明：192.168.2.1是需要改