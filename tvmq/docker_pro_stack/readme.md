
在manager上执行
docker swarm init --advertise-addr 10.105.50.201

docker network create -d overlay --attachable mongo-nw
docker network create -d overlay redis-nw
docker network create -d overlay nsq-nw
docker network create -d overlay service-nw


在manager上执行
docker stack deploy -c /opt/docker_pro_stack/nsq_stack.yml nsq-cluster
docker stack deploy -c /opt/docker_pro_stack/redis_stack.yml redis-cluster
docker stack deploy -c /opt/docker_pro_stack/mongo_stack.yml mongo-cluster
docker stack deploy -c /opt/docker_pro_stack/serv_stack.yml serv-cluster


redis-cluster配置
登录任意一个redis容器执行
docker exec -it e8622afa3860 bash
每台机器123，456，789依次排开会自动分成一主二从
redis-cli --cluster create 10.105.50.201:7101 10.105.50.201:7102 10.105.50.201:7103 10.105.219.143:7104 10.105.219.143:7105 10.105.219.143:7106 10.105.233.170:7107 10.105.233.170:7108  10.105.233.170:7109 --cluster-replicas 2


docker service ls


mongo-cluster配置

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

mongo
use admin
# 将分片加入集群

sh.addShard("shard1/shard1_1,shard1_2,shard1_3")
sh.addShard("shard2/shard2_1,shard2_2,shard2_3")
sh.addShard("shard3/shard3_1,shard3_2,shard3_3")

sh.status()


# 对数据库开启分片功能
sh.enableSharding("yaoqu")
# 对数据库中集合开启分片，并指定片键
sh.shardCollection("yaoqu.friends",{minappid:1,minopenid:1})
sh.shardCollection("yaoqu.traceinfo",{_id:1})
sh.shardCollection("yaoqu.comments",{_id:1})
sh.shardCollection("yaoqu.gifts",{_id:1})