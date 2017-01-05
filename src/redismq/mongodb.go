package redismq

import (
  "gopkg.in/mgo.v2"
)

var DB *mgo.Database

func InitDB(connStr, dbName string) *mgo.Session {
  session, err := mgo.Dial(connStr)
  if err != nil {
    panic(err)
  }
  //defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  DB = session.DB(dbName)
  return session
}

func InsertUrls(urls []string) (error) {
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

func QueryUrls(topN int) (error, []string){
  //*****查询多条数据*******
  var urls []string   //存放结果
  c := DB.C("urls")
  iter := c.Find(nil).Limit(topN).Iter()
  err := iter.All(&urls)
  return err, urls
}
