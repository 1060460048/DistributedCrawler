package redismq

import (
  "gopkg.in/mgo.v2"
)

var DB *mgo.Database

type Url struct {
  Id             bson.ObjectId `bson:"_id"`
	Url         string        `json:"url"`
}

func InitDB(connStr, dbName string) *mgo.Session {
  session, err := mgo.Dial(connStr)
  if err != nil {
    panic(err)
  }
  //defer session.Close()

  //设置模式
  session.SetMode(mgo.Monotonic, true)
  DB = session.DB(dbName)
  //获取文档集
  collection := DB.C("urls")
  // 创建索引
  index := mgo.Index{
   Key:        []string{"url"}, // 索引字段， 默认升序,若需降序在字段前加-
   Unique:     true,             // 唯一索引 同mysql唯一索引
   DropDups:   true,             // 索引重复替换旧文档,Unique为true时失效
   Background: true,             // 后台创建索引
  }
  if err := collection.EnsureIndex(index); err != nil {
     log.Println(err)
     return session
  }
  if err := collection.EnsureIndexKey("$2dsphere:location"); err != nil { // 创建一个范围索引
     log.Println(err)
     return session
  }
  return session
}

func InsertUrls(urls []Url) (error) {
  c := DB.C("urls")
  for url in urls {
    //添加去重功能
    err := c.Insert(url)
    if err != nil {
      break
    }
  }
  return err
}

func QueryUrls(topN int) (error, []Url){
  //*****查询多条数据*******
  var urls []Url   //存放结果
  c := DB.C("urls")
  iter := c.Find(nil).Limit(topN).Iter()
  err := iter.All(&urls)
  return err, urls
}
