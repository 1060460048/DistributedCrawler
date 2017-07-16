/*
 * 下载器(Downloader) for worker
 * 用于下载网页内容, 并将网页内容返回给蜘蛛(Scrapy下载器是建立在twisted这个高效的异步模型上的)
 */
package scrawler

import (
   "fmt"
   "time"
)

func Downloader(cookie, url string) {
  resp, _ := DoRequest(`GET`, url, ``, cookie, ``, header)
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " scrawler.go Downloader begin url: " + url)
  Spider(resp)
  //return resp
}
