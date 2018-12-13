一：生成image,在当前目录下运行
docker build -t rubinus/redis-cluster .

二：在manager上执行创建redis-nw
docker network create -d overlay redis-nw

三：在三台机器上分别执行
mkdir -p /opt/redis-cluster-swarm/7101/data /opt/redis-cluster-swarm/7102/data /opt/redis-cluster-swarm/7103/data /opt/redis-cluster-swarm/7104/data /opt/redis-cluster-swarm/7105/data /opt/redis-cluster-swarm/7106/data /opt/redis-cluster-swarm/7107/data /opt/redis-cluster-swarm/7108/data /opt/redis-cluster-swarm/7109/data


四：在manager上执行，进入到当前stack.yml所在的目录
docker stack deploy -c /opt/genrediscluster-swarm/stack.yml redis-cluster

docker service ls

登录任意一个容器执行
docker exec -it e8622afa3860 bash

每台机器123，456，789依次排开会自动分成一主二从
redis-cli --cluster create 10.20.80.105:7101 10.20.80.105:7102 10.20.80.105:7103 10.20.80.106:7104 10.20.80.106:7105 10.20.80.106:7106 10.20.80.132:7107 10.20.80.132:7108  10.20.80.132:7109 --cluster-replicas 2

