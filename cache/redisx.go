package cache

import (
	"encoding/json"
	"errors"

	"github.com/beego/redigo/redis"
)


// Redis cache adapter.
type RedisxCache struct {
	c        redis.Conn
	conninfo string
}

// create new redis cache with default collection name.
func NewRedisxCache() *RedisxCache {
	return &RedisxCache{}
}

// Get cache from redis.
func (rc *RedisxCache) Get(key string) interface{} {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return nil
		}
	}
	v, err := rc.c.Do("GET", key)
	if err != nil {
		return nil
	}
	return v
}

// put cache to redis.
// timeout is ignored.
func (rc *RedisxCache) Put(key string, val interface{}, timeout int64) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	
	if _, err := rc.c.Do("SET", key, val); err != nil {
		return err 
	}

	_, err := rc.c.Do("EXPIRE", key, timeout)	
	return err
}

// delete cache in redis.
func (rc *RedisxCache) Delete(key string) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	_, err := rc.c.Do("DEL", key)
	return err
}

// check cache exist in redis.
func (rc *RedisxCache) IsExist(key string) bool {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return false
		}
	}
	v, err := redis.Bool(rc.c.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// increase counter in redis.
func (rc *RedisxCache) Incr(key string) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	_, err := redis.Bool(rc.c.Do("INCRBY", key, 1))
	if err != nil {
		return err
	}
	return nil
}

// decrease counter in redis.
func (rc *RedisxCache) Decr(key string) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	_, err := redis.Bool(rc.c.Do("DECRBY", key, 1))
	if err != nil {
		return err
	}
	return nil
}

// clean all cache in redis. delete this redis collection.
func (rc *RedisxCache) ClearAll() error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	_, err := rc.c.Do("FLUSHDB")
	return err
}

// start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *RedisxCache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}
	rc.conninfo = cf["conn"]
	var err error
	rc.c, err = rc.connectInit()
	if err != nil {
		return err
	}
	if rc.c == nil {
		return errors.New("dial tcp conn error")
	}
	return nil
}

// connect to redis.
func (rc *RedisxCache) connectInit() (redis.Conn, error) {
	c, err := redis.Dial("tcp", rc.conninfo)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	Register("redisx", NewRedisxCache())
}
