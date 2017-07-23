/*
 * 项目管道(Pipeline) for worker
 * 负责处理爬虫从网页中抽取的实体，主要的功能是持久化实体、验证实体的有效性、清除不需要的信息。
 * 当页面被爬虫解析后，将被发送到项目管道，并经过几个特定的次序处理数据。
 */
package scrawler

import (
  "fmt"
  "time"
  "model"
)

func Pipeline(urls []string, items []model.Item) {

  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " pipeline.go Pipeline begin ")

  mgo := model.InitMgoDB("localhost:27017", "urls")
  defer mgo.Session.Close()
  mgo.InsertUrls(urls)
  mgo.InsertItems(items)

  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " pipeline.go Pipeline end ")
  //write urls and your data to mongodb and send finish signal to master
}
