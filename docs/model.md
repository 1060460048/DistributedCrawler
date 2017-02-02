## Distributed system for Crawl using by Golang

[![Build Status](https://travis-ci.org/zjucx/golang-webserver.svg?branch=master
)](http://120.27.39.169:8080/home)
[![Yii2](https://img.shields.io/badge/PoweredBy-ZjuCx-brightgreen.svg?style=flat)](http://120.27.39.169:8080/home)

### Introduction
使用golang开发的[分布式爬虫系统](https://github.com/zjucx/DistributedCrawler.git)，主要分为3个模块:[分布式框架](src/docs/framework.md)、[数据管理](src/docs/model.md)和[爬虫部分](src/docs/scrawler.md)。目录结构如下:
```
├── conf
│   └── app.conf       ------配置部分，数据库等信息的配置。还未开发。=。=
└── model    
    ├── mongodb.go     ------爬虫的持久化介质，存储url和想要获取的数据
    └── redismq.go     ------实用redis实现的优先级队列，master从mongodb获取url和向worker分发url
```
### [数据管理](src/docs/model.md)
```
分为持久化mongodb和内存数据库redis(实现优先级队列)。
```
#### redis
```
仅仅在master节点中使用,用作消息队列缓存要分发的url。
主要实现两个接口:
func (rmq *RedisMq) GetUrlBlock() (url string) {
  c := rmq.C

  urls, _ := redis.Strings(c.Do("brpop", "url", "0"))
  fmt.Println("url: %#v/n", urls)

  return urls[1]
}
func loadDataFromMongod() {
  //1) query 1000 urls from mongodb
  //2) push urls to redismq
}
```
#### mongodb
```
在master节点和worker节点中使用，worker节点中存储解析出来的url和最终数据;master节点中用于redis url预读。
url类数据接口:
func (mgo *Mgo)InsertUrls(urls []string) (err error) {
  c := mgo.DB.C("urls")
  for _, url := range urls {
    tmp := &Url{
      Id : bson.NewObjectId(),
      Url : url,
    }
    err = c.Insert(tmp)
    if err != nil {
      fmt.Println("insert error"+ err.Error())
      break
    }
  }
  return err
}

func (mgo *Mgo)QueryUrls(topN int) (error, []Url){
  //*****查询多条数据*******
  var urls []Url   //存放结果
  c := mgo.DB.C("urls")
  iter := c.Find(nil).Limit(topN).Iter()
  err := iter.All(&urls)
  return err, urls
}
爬取数据接口:
用户自定义的Item结构体(需要获取数据的字段)
```
### Discussing
- [submit issue](https://github.com/zjucx/DistributedCrawler/issues/new)
- email: 862575451@qq.com
