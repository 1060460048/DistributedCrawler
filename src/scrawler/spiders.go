/*
 * 爬虫(Spiders) for worker
 * 爬虫是主要干活的, 用于从特定的网页中提取自己需要的信息, 即所谓的实体(Item)。
 * 用户也可以从中提取出链接,让Scrapy继续抓取下一个页面
 */
package scrawler

import (
  "fmt"
  "time"
  "regexp"
  "strings"
  "strconv"
  "model"
)

func Spider(resp string) {
  //your code your reg you extraction rules and so on
  // 去除空格
  resp = strings.Replace(resp, " ", "", -1)
  // 去除换行符
  resp = strings.Replace(resp, "\n", "", -1)

  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " spiders.go Spider resp: " + resp)

  urls := extractUrls(resp)
  item := extractItems(resp)

  Pipeline(urls, item)
  // return item
}

func extractUrls(resp string) []string{
  var str []string

  reg := regexp.MustCompile(`<arel="next"href="(.*?)">下一页</a>`)
  arrStr := reg.FindAllStringSubmatch(resp, -1)
  if len(arrStr) > 0 {
    for i := 0; i < len(arrStr); i++ {
      str = append(str, arrStr[i][1])
      fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " spiders.go extractUrls url: " + arrStr[0][1])
    }
  }
  return str
}

func extractItems(resp string) []model.Item{
  var items []model.Item
  //resp make '\' as string so we shou add '\\' in regex and one blank become two blanks so be carefully
  //the regex should reg  the string of mstartResp
  reg := regexp.MustCompile(`<divclass="voteshidden-xs">(.*?)<small>得票</small>(.*?)">(.*?)<small>(回答|解决)</small>(.*?)<span>(.*?)</span><small>浏览</small>(.*?)<h2class="title"><ahref="(.*?)">(.*?)</a></h2>`)
  arrStr := reg.FindAllStringSubmatch(resp, -1)
  if len(arrStr) > 0 {
    for i := 0; i < len(arrStr); i++ {
      var item model.Item
      item.Votes, _ = strconv.Atoi(arrStr[i][1])
      item.Answers, _ = strconv.Atoi(arrStr[i][3])
      item.Views, _ = strconv.Atoi(arrStr[i][6])
      item.Url = arrStr[i][6]
      item.Question = arrStr[i][9]
      items = append(items, item)
      fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " spiders.go extractUrls url: "+ arrStr[i][1] + arrStr[i][3] + arrStr[i][6] +arrStr[i][8] + arrStr[i][9])
    }
  }
  return items
}
