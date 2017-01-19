/*
 * 项目管道(Pipeline) for worker
 * 负责处理爬虫从网页中抽取的实体，主要的功能是持久化实体、验证实体的有效性、清除不需要的信息。
 * 当页面被爬虫解析后，将被发送到项目管道，并经过几个特定的次序处理数据。
 */
package scrawler

import (
  "fmt"
  "model"
)


// type Item struct {
//   name string
//   sex  string
//   habbit []string
//   //define by yourself...
// }

func Pipeline(urls, items []string) {
  fmt.Printf("Pipeline======")
  mgo := model.InitMgoDB()
  mgo.InsertUrls(urls)
  defer mgo.MgoClient.Close()
  //write urls and your data to mongodb and send finish signal to master
}
