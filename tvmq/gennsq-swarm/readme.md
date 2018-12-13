在manager上执行

docker network create -d overlay nsq-nw

在三台机器上分别执行
mkdir -p /opt/nsq-cluster-swarm

在manager上执行
docker stack deploy -c /opt/gennsq-swarm/stack.yml nsq-cluster

docker service ls
