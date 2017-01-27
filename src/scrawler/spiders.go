/*
 * 爬虫(Spiders) for worker
 * 爬虫是主要干活的, 用于从特定的网页中提取自己需要的信息, 即所谓的实体(Item)。
 * 用户也可以从中提取出链接,让Scrapy继续抓取下一个页面
 */
package scrawler

import (
  // "fmt"
  "regexp"
  // "strings"
  "model"
)

func Spider(resp string) {
  //your code your reg you extraction rules and so on
  urls := extractUrls(resp)
  item := extractItems(resp)

  Pipeline(urls, item)
  // return item
}

func extractUrls(resp string) []string{
  var str []string
  return str
}

func extractItems(resp string) *model.Item{
  item := &model.Item{}
  //resp make '\' as string so we shou add '\\' in regex and one blank become two blanks so be carefully
  //the regex should reg  the string of mstartResp
  reg := regexp.MustCompile(`<a class=\\"S_txt1\\" target=\\"_blank\\"  usercard=\\"(.*?)\\" href=\\"(.*?)\\" title=\\"(.*?)\\">`)
  /*arrStart := */reg.FindAllStringSubmatch(resp, -1)
  //
  // if len(arrStart) > 0 {
  //   for i := 0; i < len(arrStart); i++ {
  //     mUrl[arrStart[i][3]] = strings.Replace(strings.Replace(arrStart[i][2], "\\/", "/", -1), "weibo.com", "weibo.cn", -1)
  //     fmt.Println(mUrl[arrStart[i][3]])
  //     //item.Proper = mUrl[arrStart[i][3]]
  //   }
  // }
  return item
}
