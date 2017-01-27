package scrawler

import (
  "fmt"
)

// func Scrawler(cookie string, urls []string){
func Scrawler(){
  // get login cookies
  //cookie := WeiboLogin("hfutcx@163.com", "dao88xiang")
  urls, err := ReadLine("src/data/mstarturlname.map")
  if err != nil {
    fmt.Println("ReadLine error")
    return
  }
  Scheduler("cookie", urls)
}
