package scrawler

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
  // "os"
  // "bufio"
  // "path/filepath"
)

// func Scrawler(cookie string, urls []string){
func Scrawler(){
  // get login cookies
  cookie := WeiboLogin("username", "passwd")
  urls, err := ReadLine("src/data/mstarturlname.map")
  if err != nil {
    fmt.Println("ReadLine error")
    return
  }
  Scheduler(cookie, urls)
}

var mstartUrl = "http://d.weibo.com/1087030002_2975_1003_0"

func ScrawlerBak(username, passwd string){
  // get login cookies
  loginCookies := WeiboLogin(username, passwd)

  mUrl := make(map[string]string)
  for i := 1; i < 151; i++ {
    // mstartUrl = "http://d.weibo.com/1087030002_2975_1003_0?pids=Pl_Core_F4RightUserList__4&page=" + strconv.Itoa(i) + "#Pl_Core_F4RightUserList__4"
    mstartUrl = "http://d.weibo.com/1087030002_2975_2006_0?page=" + strconv.Itoa(i) + "#Pl_Core_F4RightUserList__4"
    mstartResp, _ := DoRequest(`GET`, mstartUrl, ``, loginCookies, ``, header)
    //fmt.Println(mstartResp)
    //resp make '\' as string so we shou add '\\' in regex and one blank become two blanks so be carefully
    //the regex should reg  the string of mstartResp
    reg := regexp.MustCompile(`<a class=\\"S_txt1\\" target=\\"_blank\\"  usercard=\\"(.*?)\\" href=\\"(.*?)\\" title=\\"(.*?)\\">`)
    arrStart := reg.FindAllStringSubmatch(mstartResp, -1)

    if len(arrStart) > 0 {
      for i := 0; i < len(arrStart); i++ {
        mUrl[arrStart[i][3]] = strings.Replace(strings.Replace(arrStart[i][2], "\\/", "/", -1), "weibo.com", "weibo.cn", -1)
        fmt.Println(mUrl[arrStart[i][3]])
      }
    }
  }

  // for k, v := range mUrl {
  //   getPageData("./data/" + k, v, loginCookies)
  // }
  // writeMaptoFile(mUrl, "./data/mstarturlname.map")
}
