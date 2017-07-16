package scrawler

import (
   "fmt"
   "time"
)

// func Scrawler(cookie string, urls []string){
func Scrawler(url string) error {
  // get login cookies
  //cookie := WeiboLogin("hfutcx@163.com", "xxxxxx")
  //urls, err := ReadLine("src/data/mstarturlname.map")
  //if err != nil {
  //  fmt.Println("ReadLine error")
  //  return
  //}
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " scrawler.go Scrawler begin url: " + url)
  Downloader("cookie", url)
  return nil
}
