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
  Mgo         *Mgo
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
  rmq.Mgo = InitMgoDB("localhost:27017", "urls")
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
    urls := rmq.GetUrls()
    // if len(urls) == 0 {
    //   fmt.Println("loadUrlsFromRedis urls is nil sleep 60s")
    //   time.Sleep(60 * time.Second)
    // }
    fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadUrlsFromRedis add " + urls[1] + " to jobChan")
    jobChan <- urls[1]
  }
}

func (rmq *RedisMq) GetUrls() (urls []string) {
  c := rmq.C
  // defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  n, _ := redis.Int(c.Do("LLEN", "url"))

  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrls url num: %d", n)
  // if len(urls) < 100 then load data from mongodb
  if n < 100 {
    rmq.loadDataFromMongod(100-n)
  }
  urls, _ = redis.Strings(c.Do("brpop", "url", "0"))
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go GetUrls url: " + urls[1])
  return urls
}

func (rmq *RedisMq) PushUrls(urls []Url) {
  // rc := rmq.RedisClient.Get()
  rc := rmq.C
  // defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  for _, url := range urls {
    rc.Do("lpush", "url", url.Url)
  }
}

func (rmq *RedisMq) PushUrl(url Url) {
  // rc := rmq.RedisClient.Get()
  rc := rmq.C
  // defer rc.Close()
  //values, _ := redis.Values(rc.Do("lrange", "redlist", "0", "100")))
  rc.Do("lpush", "url", url.Url)
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go PushUrl url: " + url.Url)
}

func (rmq *RedisMq) loadDataFromMongod(n int) {
  //1) queru 1000 urls from mongodb
  //2) push urls to redismq
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " redismq.go loadDataFromMongod ")
  urls, _:= rmq.Mgo.QueryUrls(n)
  for _, url := range urls {
    rmq.PushUrl(url)
    rmq.Mgo.DeleteUrl(url)
  }
}
