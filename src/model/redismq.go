package model

import (
  "fmt"
  "time"
  "github.com/garyburd/redigo/redis"
)

type RedisMq struct {
  RedisClient *redis.Pool
  redisHost   string
  redisDB     int
  C           redis.Conn
}

func InitRedisMq(redisHost string, redisDB int) (*RedisMq, error) {
  rmq := &RedisMq{
    redisHost : redisHost,
    redisDB : redisDB,
  }
  c, err := redis.Dial("tcp", redisHost)
  if err != nil {
      // handle error
      fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go InitRedisMq error: " + err.Error())
      return nil, err
  }
  c.Do("SELECT", rmq.redisDB)
  rmq.C = c
  return rmq, nil
}

/*
 * this function is likely a producter.
 */
func (rmq *RedisMq) LoadUrlsFromRedis(jobChan chan string) {
  //1) load Data
  //2) dispatchjob
  // When finish you need dispatchjob for
  // every blocked work because of none data in redis
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis begin")
  for {
    url := rmq.GetUrlBlock()
    // if len(urls) == 0 {
    //   fmt.Println("loadUrlsFromRedis urls is nil sleep 60s")
    //   time.Sleep(60 * time.Second)
    // }
    fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis add " + url + " to jobChan")
    jobChan <- url
  }
}

func (rmq *RedisMq) GetUrlBlock() (url string) {
  c := rmq.C

  urls, _ := redis.Strings(c.Do("brpop", "url", "0"))
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrlBlock url: " + urls[1])
  return urls[1]
}

func (rmq *RedisMq) GetUrls() (urls []string) {
  c := rmq.C
  // defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  n, _ := redis.Int(c.Do("LLEN", "url"))

  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrls url num: %#v", n)
  // if len(urls) < 100 then load data from mongodb
  if n < 100 {
    go loadDataFromMongod()
  } else {
    n = 100
  }
  for ;n > 0; n-- {
    url, _ := redis.Strings(c.Do("brpop", "url", "0"))
    fmt.Println("url: %#v/n", url)
    //urls = append(urls, url)
  }
  return urls
}

func (rmq *RedisMq) PushUrls(urls []string) {
  // rc := rmq.RedisClient.Get()
  rc := rmq.C
  // defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  // for url := l.Front; url != nil; url = url.Next() {
  rc.Do("lpush", "url", urls)
  // }
}

func loadDataFromMongod() {
  //1) queru 1000 urls from mongodb
  //2) push urls to redismq
}
