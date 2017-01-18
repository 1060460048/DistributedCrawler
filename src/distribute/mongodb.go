package distribute

import (
  "gopkg.in/mgo.v2"
)

type Mgo struct {
  MgoClient *mgo.Session
  MgoHost string
  MgoDB int
}

type Url struct {
  Id             bson.ObjectId `bson:"_id"`
	Url            string        `json:"url"`
}

func InitMgoDB(ConnStr, DBName string) *mgo.Session {
  mgo := &Mgo{
    MgoHost : ConnStr
    MgoDB : DBName
  }
  session, err := mgo.Dial(ConnStr)
  if err != nil {
    panic(err)
  }
  //defer session.Close()

  //设置模式
  session.SetMode(mgo.Monotonic, true)
  //获取文档集
  collection = session.DB(DBName).C("urls")
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
  mgo.MgoClient = session
  return session
}

func (mgo *Mgo)InsertUrls(urls []string) (error) {
  c := mgo.MgoClient
  for url in urls {
    tmp := &Url{
      Url : url
    }
    err := c.Insert(tmp)
    if err != nil {
      break
    }
  }
  return err
}

func (mgo *Mgo)QueryUrls(topN int) (error, []Url){
  c := mgo.MgoClient
  //*****查询多条数据*******
  var urls []Url   //存放结果
  c := DB.C("urls")
  iter := c.Find(nil).Limit(topN).Iter()
  err := iter.All(&urls)
  return err, urls
}
