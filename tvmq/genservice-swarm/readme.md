在manager上执行

docker network create -d overlay service-nw

在manager上执行
docker stack deploy -c /opt/genservice-swarm/stack.yml service-cluster

docker service ls
