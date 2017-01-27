package model

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Mgo struct {
  Session *mgo.Session
  DB      *mgo.Database
  MgoHost string
  MgoDB   string
}

type Url struct {
  Id             bson.ObjectId `bson:"_id"`
	Url            string        `json:"url"`
}

type Item struct {
  id              bson.ObjectId `bson:"_id"`
  name            string        `bson:"name"`
  sex             string        `bson:"sex"`
  habbit          []string      `bson:"habbit"`
  //define by yourself...
}

func InitMgoDB(ConnStr, DBName string) *Mgo {
  mgoClient := &Mgo{
    MgoHost : ConnStr,
    MgoDB : DBName,
  }
  session, err := mgo.Dial(ConnStr)
  if err != nil {
    panic(err)
  }
  //defer session.Close()

  //设置模式
  session.SetMode(mgo.Monotonic, true)
  //获取文档集
  c := session.DB(DBName).C("urls")
  // 创建索引
  index := mgo.Index{
   Key:        []string{"url"}, // 索引字段， 默认升序,若需降序在字段前加-
   Unique:     true,             // 唯一索引 同mysql唯一索引
   DropDups:   true,             // 索引重复替换旧文档,Unique为true时失效
   Background: true,             // 后台创建索引
  }
  if err := c.EnsureIndex(index); err != nil {
     fmt.Println(err)
     return nil
  }
  mgoClient.DB = session.DB(DBName)
  mgoClient.Session = session
  return mgoClient
}

func (mgo *Mgo)Close(){
  mgo.Session.Close()
}

func (mgo *Mgo)InsertUrls(urls []string) (err error) {
  c := mgo.DB.C("urls")
  for _, url := range urls {
    tmp := &Url{
      Url : url,
    }
    err = c.Insert(tmp)
    if err != nil {
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
