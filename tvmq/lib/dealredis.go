package lib

import (
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib/rdcluster"
	"code.tvmining.com/tvplay/tvmq/lib/rdone"
	"code.tvmining.com/tvplay/tvmq/redisclusterpool"
	"code.tvmining.com/tvplay/tvmq/redispool"
)

func JudgeScard(key string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Scard(redisclusterpool.GetConn(), key)
	} else {
		return rdone.Scard(redispool.GetConn(), key)
	}
}

func JudgeHmsetAndSet(key string, fvalues interface{},
	setKey, member string, ttl string) (string, int64) {
	if config.UseRedisCluster == 1 {
		return rdcluster.HmsetAndSet(redisclusterpool.GetConn(),
			key, fvalues, redisclusterpool.GetConn(), setKey, member, ttl)
	} else {
		return rdone.HmsetAndSet(redispool.GetConn(), key, fvalues, redispool.GetConn(), setKey, member, ttl)
	}
}

func JudgeHmsetAndHmset(key string, fvalues interface{},
	key2 string, values interface{}, ttl string) (string, string) {
	if config.UseRedisCluster == 1 {
		return rdcluster.HmsetAndHmset(redisclusterpool.GetConn(),
			key, fvalues, redisclusterpool.GetConn(), key2, values, ttl)
	} else {
		return rdone.HmsetAndHmset(redispool.GetConn(), key, fvalues, redispool.GetConn(), key2, values, ttl)
	}
}

func JudgeHmset(key string, fvalues interface{}, ttl string) (string, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Hmset(redisclusterpool.GetConn(), key, fvalues)
	} else {

		return rdone.Hmset(redispool.GetConn(), key, fvalues, ttl)
	}
}

func JudgeSadd(key, member string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Sadd(redisclusterpool.GetConn(), key, member)
	} else {

		return rdone.Sadd(redispool.GetConn(), key, member)
	}
}

func JudgeSmembers(key string) ([]string, error) {
	var reply []string
	if config.UseRedisCluster == 1 {
		return rdcluster.Smembers(redisclusterpool.GetConn(), key, reply)
	} else {

		return rdone.Smembers(redispool.GetConn(), key, reply)
	}
}

func JudgeHgetall(key string) (map[string]string, error) {
	m := make(map[string]string)
	if config.UseRedisCluster == 1 {
		return rdcluster.Hgetall(redisclusterpool.GetConn(), key, m)
	} else {
		return rdone.Hgetall(redispool.GetConn(), key, m)
	}
	//mapstructure.Decode(m, &v)
}

func JudgeSetex(key string, ttl int, value string) (string, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Setex(redisclusterpool.GetConn(), key, ttl, value)
	} else {

		return rdone.Setex(redispool.GetConn(), key, ttl, value)
	}
}

func JudgeGet(key string) (string, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Get(redisclusterpool.GetConn(), key)
	} else {

		return rdone.Get(redispool.GetConn(), key)
	}
}

func JudgeSrem(key, member string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Srem(redisclusterpool.GetConn(), key, member)
	} else {

		return rdone.Srem(redispool.GetConn(), key, member)
	}
}

func JudgeDelKey(key string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.DelKey(redisclusterpool.GetConn(), key)
	} else {

		return rdone.DelKey(redispool.GetConn(), key)
	}
}

func JudgeDelKeyAndSetMember(key, setKey, member string) (int64, int64) {
	if config.UseRedisCluster == 1 {
		return rdcluster.DelKeyAndSetMember(redisclusterpool.GetConn(), key,
			redisclusterpool.GetConn(), setKey, member)
	} else {

		return rdone.DelKeyAndSetMember(redispool.GetConn(), key,
			redispool.GetConn(), setKey, member)
	}
}

func JudgeZadd(key, score, member string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Zadd(redisclusterpool.GetConn(), key, score, member)
	} else {

		return rdone.Zadd(redispool.GetConn(), key, score, member)
	}
}

func JudgeZrange(key, start, stop string, st int) ([]string, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Zrange(redisclusterpool.GetConn(), key, start, stop, st)
	} else {

		return rdone.Zrange(redispool.GetConn(), key, start, stop, st)
	}
}

func JudgeZcard(key string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Zcard(redisclusterpool.GetConn(), key)
	} else {
		return rdone.Zcard(redispool.GetConn(), key)
	}
}

func JudgeZremrangebyrank(key, start, stop string) (int64, error) {
	if config.UseRedisCluster == 1 {
		return rdcluster.Zremrangebyrank(redisclusterpool.GetConn(), key, start, stop)
	} else {
		return rdone.Zremrangebyrank(redispool.GetConn(), key, start, stop)
	}
}
