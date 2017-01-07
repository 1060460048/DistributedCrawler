package redismq

import (
  "fmt"
  //"sync"
  //"net"
  //"net/rpc"
  "container/list"
  "time"
  "github.com/garyburd/redigo/redis"
)

type RedisMq struct {
  RedisClient *redis.Pool
  RedisHost string
  RedisDB int
}

func initRedisMq(RedisHost string, RedisDB int) *RedisMq {
  rmq := &RedisMq{
    RedisHost : RedisHost,
    RedisDB : RedisDB,
  }
  rmq.RedisClient = &redis.Pool{
      MaxIdle:     1,
  		MaxActive:   10,
  		IdleTimeout: 180 * time.Second,
  		Dial: func() (redis.Conn, error) {
  			c, err := redis.Dial("tcp", rmq.RedisHost)
  			if err != nil {
  				return nil, err
  			}
  			// 选择db
  			c.Do("SELECT", rmq.RedisDB)
  			return c, nil
  		},
    }
    return rmq
}

func RunRedisMq(RedisHost string, RedisDB int) {
  rmq := initRedisMq(RedisHost, RedisDB)
  t := time.NewTicker(60 * time.Second)
  fmt.Println("RunRedisMq: ", RedisHost, " RedisDB: ", RedisDB)
  for {
    select {
    case <-t.C:
      readUrlFromMongod(rmq)
    }
  }
}

func (rmq *RedisMq) GetUrls() []string{
  rc := rmq.RedisClient.Get()
  defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  urls, _ := redis.Strings(c.Do("lrange", "url", "0", "100"))
  fmt.Printf("get urls from redis: " + urls)
  // if len(urls) < 100 then load data from mongodb
  // loadDataFromMongod()
  return urls
}

func (rmq *RedisMq) PushUrls(urls []string) {
  rc := rmq.RedisClient.Get()
  defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  // for url := l.Front; url != nil; url = url.Next() {
  rc.Do("lpush", "url", urls)
  // }
}

func (rmq *RedisMq) LoadDataFromMongod(lengthName string) {
  
}
