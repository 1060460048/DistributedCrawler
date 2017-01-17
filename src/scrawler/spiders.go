/*
 * 爬虫(Spiders)
 * 爬虫是主要干活的, 用于从特定的网页中提取自己需要的信息, 即所谓的实体(Item)。
 * 用户也可以从中提取出链接,让Scrapy继续抓取下一个页面
 */
 package scrawler

 import (
   "fmt"
   "regexp"
   "strconv"
   "strings"
   "os"
   "bufio"
   "path/filepath"
 )
