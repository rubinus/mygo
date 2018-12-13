在manager上执行
docker swarm init --advertise-addr  192.168.30.73
docker network create -d overlay --attachable mongo-nw

在三台机器上分别执行
mkdir -p /opt/mongo-cluster-swarm/configsvr /opt/mongo-cluster-swarm/shard1 /opt/mongo-cluster-swarm/shard2 /opt/mongo-cluster-swarm/shard3

docker pull mongo


在manager上执行
docker stack deploy -c /opt/genmongocluster-swarm/stack.yml mongo-cluster

docker service ls


进入shard1副本集：
# 宿主机
docker exec -it mongo_shard1_1.1***** bash
# 容器中
mongo
切换到admin库
use admin

rs.initiate({
    "_id":"shard1",
    "members":[
        {
            "_id":0,
            "host":"shard1_1"
        },
        {
            "_id":1,
            "host":"shard1_2"
        },
        {
            "_id":2,
            "host":"shard1_3",
            "arbiterOnly":true
        }
    ],
    "settings":{
        "getLastErrorDefaults" :
            {
                "w" : "majority",
                "wtimeout" : 30000
            }
    }
})
rs.status()


进入shard2副本集：
# 宿主机
docker exec -it mongo_shard2_1.1**** bash
# 容器中
mongo
切换到admin库
use admin

rs.initiate({
    "_id":"shard2",
    "members":[
        {
            "_id":0,
            "host":"shard2_1"
        },
        {
            "_id":1,
            "host":"shard2_2",
            "arbiterOnly":true
        },
        {
            "_id":2,
            "host":"shard2_3"
        }
    ],
    "settings":{
        "getLastErrorDefaults" :
            {
                "w" : "majority",
                "wtimeout" : 30000
            }
    }
})
rs.status()


进入shard3副本集：
# 宿主机
docker exec -it mongo_shard3_1.1**** bash
# 容器中
mongo
切换到admin库
use admin

rs.initiate({
    "_id":"shard3",
    "members":[
        {
            "_id":0,
            "host":"shard3_1",
            "arbiterOnly":true
        },
        {
            "_id":1,
            "host":"shard3_2"
        },
        {
            "_id":2,
            "host":"shard3_3"
        }
    ],
    "settings":{
        "getLastErrorDefaults" :
            {
                "w" : "majority",
                "wtimeout" : 30000
            }
    }
})
rs.status()


进入cfg_1副本集：
# 宿主机
docker exec -it cfg_1**** bash
# 容器中
mongo
切换到admin库
use admin

rs.initiate({
    "_id":"cfg",
    "configsvr":true,
    "members":[
        {
            "_id":0,
            "host":"cfg_1"
        },
        {
            "_id":1,
            "host":"cfg_2"
        },
        {
            "_id":2,
            "host":"cfg_3"
        }
    ]
})
rs.status()



# 连接mongos，端口号与mongos配置文件中设定一致

mongo -port 22000
use admin
# 将分片加入集群

sh.addShard("shard1/shard1_1,shard1_2,shard1_3")
sh.addShard("shard2/shard2_1,shard2_2,shard2_3")
sh.addShard("shard3/shard3_1,shard3_2,shard3_3")

sh.status()